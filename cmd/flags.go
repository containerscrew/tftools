package cmd

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(summarizeCmd)
	rootCmd.AddCommand(usageCmd)
}
