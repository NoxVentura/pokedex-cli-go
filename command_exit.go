package main

func commandExit(args ...string) error {
	isStopped = true
	return nil
}
