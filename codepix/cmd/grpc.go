package cmd

import (
	"os"

	"github.com/FabioMeirelles3/codepix/application/grpc"
	"github.com/FabioMeirelles3/codepix/infrastructure/db"
	"github.com/spf13/cobra"
)

var portNumber int

// grpcCmd represents the grpc command
var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "Start gRPC server",
	Run: func(cmd *cobra.Command, args []string) {
		database := db.ConnectDB(os.Getenv("env"))
		grpc.StartGrpcServer(database, portNumber)
	},
}

func init() {
	rootCmd.AddCommand(grpcCmd)
	grpcCmd.Flags().IntVarP(&portNumber, "port", "p", 50051, "gRPC Server Port")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// grpcCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// grpcCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
