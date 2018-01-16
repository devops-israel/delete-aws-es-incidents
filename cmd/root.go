// Copyright © 2017 Josh Dvir <josh@dvir.uk>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"strings"

	cron "github.com/robfig/cron"
	"github.com/spf13/cobra"
	elastic "gopkg.in/olivere/elastic.v5"
)

var (
	olderThanInDays int
	esURL           string
	prefixes        string
	ctx             context.Context
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "delete-aws-es-incidents",
	Short: "Delete ELK incidents on AWS ES 5.1",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		if esURL == "" {
			println("No Elasticsearch URL present, can't continue.")
			os.Exit(0)
		}

		client, err := elastic.NewClient(
			elastic.SetURL(esURL),
			elastic.SetSniff(false),
		)
		if err != nil {
			panic(err)
		}

		esversion, err := client.ElasticsearchVersion(esURL)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Elasticsearch version %s\n", esversion)

		runCommand()
		var wgm sync.WaitGroup
		cron := cron.New()
		cron.AddFunc("@hourly", func() { runCommand() })
		cron.Start()
		println("Cron run started...")
		wgm.Add(1)
		wgm.Wait()
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.Flags().IntVarP(&olderThanInDays, "older-than-in-days", "d", 14, "delete incidents older then in days")
	RootCmd.Flags().StringVarP(&esURL, "es-url", "e", "", "Elasticsearch URL, eg. https://path-to-es.aws.com/")
	RootCmd.Flags().StringVarP(&prefixes, "prefixes", "p", "logstash-", "comma separated list of prefixes for indexs, index date must be in format YYYY.MM.DD. eg. 'logstash-2017.09.28'. default is 'logstash-'")
}

func runCommand() {
	println("Starting deleting incidents run...")
	ctx = context.Background()
	client, err := elastic.NewClient(
		elastic.SetURL(esURL),
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}

	prefixesArr := strings.Split(prefixes, ",")
	for _, prefix := range prefixesArr {
		indexNames, err := client.IndexNames()
		if err != nil {
			panic(err)
		}

		for _, indexName := range indexNames {
			if strings.HasPrefix(indexName, prefix) {
				date := strings.TrimPrefix(indexName, prefix)
				dateArr := strings.Split(date, ".")
				nowTime := time.Now()
				indexYear, _ := strconv.Atoi(dateArr[0])
				indexMonth, _ := strconv.Atoi(dateArr[1])
				indexDay, _ := strconv.Atoi(dateArr[2])
				incidentTime := time.Date(indexYear, time.Month(indexMonth), indexDay, 0, 0, 0, 0, nowTime.Location())
				fmt.Printf("daysDiff %d\n", daysDiff(nowTime, incidentTime))
				if daysDiff(nowTime, incidentTime) > olderThanInDays {
					fmt.Printf("About to delete index %s\n", indexName)
					if (prefix == "logstash-"){
						deleteIncident(ctx, client, indexName)
					}
				}
			}
		}
	}
	println("Ending deleting incidents run...")
}

func deleteIncident(ctx context.Context, client *elastic.Client, indexName string) {
	deleteIndex, err := client.DeleteIndex(indexName).Do(ctx)
	if err != nil {
		fmt.Printf("Error deleting index %s\n", indexName)
	}

	if deleteIndex.Acknowledged {
		fmt.Printf("index %s deleted.\n", indexName)
	}
}

func lastDayOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), 12, 31, 0, 0, 0, 0, t.Location())
}

func firstDayOfNextYear(t time.Time) time.Time {
	return time.Date(t.Year()+1, 1, 1, 0, 0, 0, 0, t.Location())
}

// a - b in days
func daysDiff(a, b time.Time) (days int) {
	cur := b
	for cur.Year() < a.Year() {
		// add 1 to count the last day of the year too.
		days += lastDayOfYear(cur).YearDay() - cur.YearDay() + 1
		cur = firstDayOfNextYear(cur)
	}
	days += a.YearDay() - cur.YearDay()
	if b.AddDate(0, 0, days).After(a) {
		days--
	}
	return days
}
