package main

func main() {
	driver := &Driver{
		endpoint: "unix:///csi/csi.sock",
		nodeID:   "1",
	}
	driver.Run()
}
