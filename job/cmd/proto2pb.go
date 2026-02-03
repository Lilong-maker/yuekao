/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// proto2pbCmd represents the proto2pb command
var proto2pbCmd = &cobra.Command{
	Use:   "proto2pb",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		exec.Command(
			"protoc",
			"-I", ".",
			"--go_out=pb",
			"--go-grpc_out==pb",
			"goods.proto",
		).Run()
		fmt.Println("proto2pb called")
	},
}

func init() {
	rootCmd.AddCommand(proto2pbCmd)
	//dbType := flag.String("db", "mysql", "the database type")
	//host := flag.String("host", "localhost", "the database host")
	//port := flag.Int("port", 3306, "the database port")
	//user := flag.String("user", "root", "the database user")
	//password := flag.String("password", "", "the database password")
	//schema := flag.String("schema", "", "the database schema")
	//table := flag.String("table", "*", "the table schema，multiple tables ',' split. ")
	//serviceName := flag.String("service_name", *schema, "the protobuf service name , defaults to the database schema.")
	//packageName := flag.String("package", *schema, "the protocol buffer package. defaults to the database schema.")
	//goPackageName := flag.String("go_package", "", "the protocol buffer go_package. defaults to the database schema.")
	//ignoreTableStr := flag.String("ignore_tables", "", "a comma spaced list of tables to ignore")
	//ignoreColumnStr := flag.String("ignore_columns", "", "a comma spaced list of mysql columns to ignore")
	//fieldStyle := flag.String("field_style", "sqlPb", "gen protobuf field style, sql_pb | sqlPb")
	////Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// proto2pbCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// proto2pbCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
