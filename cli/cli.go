package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	iggy "github.com/eldpcn/iggy-go"
)

// CLI commands
var (
	createStreamCmd = flag.NewFlagSet("createstream", flag.ExitOnError)
	getStreamCmd    = flag.NewFlagSet("getstream", flag.ExitOnError)
	deleteStreamCmd = flag.NewFlagSet("deletestream", flag.ExitOnError)

	createTopicCmd = flag.NewFlagSet("createtopic", flag.ExitOnError)
	getTopicCmd    = flag.NewFlagSet("gettopic", flag.ExitOnError)
	deleteTopicCmd = flag.NewFlagSet("deletetopic", flag.ExitOnError)
)

// CLI flags
var (
	url  string
	port string

	//get stream
	gs_streamId int

	//create stream
	cs_streamId int
	cs_name     string

	//delete stream
	ds_streamId int

	//get topic
	gt_streamId int
	gt_topicId  int

	//create topic
	ct_streamId        int
	ct_topicId         int
	ct_name            string
	ct_partitionsCount int

	//delete topic
	dt_streamId int
	dt_topicId  int
)

func getUsage() {
	fmt.Println("Usage:")
	fmt.Println("  getstream -url <url> -port <port> -streamId <streamId>")
	fmt.Println("  createstream -url <url> -port <port> -streamId <streamId> -name <name>")
	fmt.Println("  deletestream -url <url> -port <port> -streamId <streamId>")
	fmt.Println("  gettopic -url <url> -port <port> -streamId <streamId> -topicId <topicId>")
	fmt.Println("  createtopic -url <url> -port <port> -streamId <streamId> -topicId <topicId> -name <name> -partitionsCount <partitionsCount>")
	fmt.Println("  deletetopic -url <url> -port <port> -streamId <streamId> -topicId <topicId>")
}

func init() {
	getStreamCmd.StringVar(&url, "url", "127.0.0.1", "Iggy server url")
	getStreamCmd.StringVar(&port, "port", "8090", "Iggy server port")
	getStreamCmd.IntVar(&gs_streamId, "streamid", -1, "Stream Id. Returns all streams for default value.")
	getStreamCmd.IntVar(&gs_streamId, "sid", -1, "Alias for Stream Id")
	getStreamCmd.IntVar(&gs_streamId, "s", -1, "Alias for Stream Id")

	createStreamCmd.StringVar(&url, "url", "127.0.0.1", "Iggy server url")
	createStreamCmd.StringVar(&port, "port", "8090", "Iggy server port")
	createStreamCmd.IntVar(&cs_streamId, "streamid", 1, "Stream Id")
	createStreamCmd.IntVar(&cs_streamId, "sid", 1, "Alias for Stream Id")
	createStreamCmd.IntVar(&cs_streamId, "s", 1, "Alias for Stream Id")
	createStreamCmd.StringVar(&cs_name, "name", "", "Stream name")
	createStreamCmd.StringVar(&cs_name, "n", "", "Stream name")

	deleteStreamCmd.StringVar(&url, "url", "127.0.0.1", "Iggy server url")
	deleteStreamCmd.StringVar(&port, "port", "8090", "Iggy server port")
	deleteStreamCmd.IntVar(&ds_streamId, "streamid", -1, "Stream Id")
	deleteStreamCmd.IntVar(&ds_streamId, "sid", -1, "Alias for Stream Id")
	deleteStreamCmd.IntVar(&ds_streamId, "s", -1, "Alias for Stream Id")

	getTopicCmd.StringVar(&url, "url", "127.0.0.1", "Iggy server url")
	getTopicCmd.StringVar(&port, "port", "8090", "Iggy server port")
	getTopicCmd.IntVar(&gt_streamId, "streamid", 1, "Stream Id")
	getTopicCmd.IntVar(&gt_streamId, "sid", 1, "Alias for Stream Id")
	getTopicCmd.IntVar(&gt_streamId, "s", 1, "Alias for Stream Id")
	getTopicCmd.IntVar(&gt_topicId, "topicid", -1, "Topic Id. Returns all streams for default value.")
	getTopicCmd.IntVar(&gt_topicId, "tid", -1, "Alias for Topic Id")
	getTopicCmd.IntVar(&gt_topicId, "t", -1, "Alias for Topic Id")

	createTopicCmd.StringVar(&url, "url", "127.0.0.1", "Iggy server url")
	createTopicCmd.StringVar(&port, "port", "8090", "Iggy server port")
	createTopicCmd.IntVar(&ct_streamId, "streamid", 1, "Stream Id")
	createTopicCmd.IntVar(&ct_streamId, "sid", 1, "Alias for Stream Id")
	createTopicCmd.IntVar(&ct_streamId, "s", 1, "Alias for Stream Id")
	createTopicCmd.IntVar(&ct_topicId, "topicid", 2, "Topic Id")
	createTopicCmd.IntVar(&ct_topicId, "tid", 1, "Alias for Topic Id")
	createTopicCmd.IntVar(&ct_topicId, "t", 1, "Alias for Topic Id")
	createTopicCmd.StringVar(&ct_name, "name", "", "Topic name")
	createTopicCmd.StringVar(&ct_name, "n", "", "Alias for Topic name")
	createTopicCmd.IntVar(&ct_partitionsCount, "partitionsCount", 1, "Number of partitions in topic")
	createTopicCmd.IntVar(&ct_partitionsCount, "pc", 1, "Alias for Number of partitions in topic")

	deleteTopicCmd.StringVar(&url, "url", "127.0.0.1", "Iggy server url")
	deleteTopicCmd.StringVar(&port, "port", "8090", "Iggy server port")
	deleteTopicCmd.IntVar(&dt_streamId, "streamid", -1, "Stream Id")
	deleteTopicCmd.IntVar(&dt_streamId, "sid", -1, "Alias for Stream Id")
	deleteTopicCmd.IntVar(&dt_streamId, "s", -1, "Alias for Stream Id")
	deleteTopicCmd.IntVar(&dt_topicId, "topicid", -1, "Topic Id")
	deleteTopicCmd.IntVar(&dt_topicId, "tid", -1, "Alias for Topic Id")
	deleteTopicCmd.IntVar(&dt_topicId, "t", -1, "Alias for Topic Id")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: No command provided.")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "createstream":
		createStreamCmd.Parse(os.Args[2:])
		ms := CreateMessageStream()
		if cs_name == "" {
			fmt.Println("Error: Name flag is required.")
			createStreamCmd.PrintDefaults()
			os.Exit(1)
		}

		err := ms.CreateStream(iggy.StreamRequest{
			StreamId: cs_streamId,
			Name:     cs_name,
		})
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

	case "getstream":
		getStreamCmd.Parse(os.Args[2:])
		ms := CreateMessageStream()
		if gs_streamId == -1 {
			streams, err := ms.GetStreams()
			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}
			jsonData, err := json.MarshalIndent(&streams, "", "  ")
			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}
			fmt.Println(string(jsonData))
			return
		}

		stream, err := ms.GetStreamById(gs_streamId)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		jsonData, err := json.MarshalIndent(&stream, "", "  ")
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println(string(jsonData))

	case "deletestream":
		deleteStreamCmd.Parse(os.Args[2:])
		ms := CreateMessageStream()
		if ds_streamId == -1 {
			fmt.Println("Error: Stream Id is required.")
			deleteStreamCmd.PrintDefaults()
			os.Exit(1)
		}

		err := ms.DeleteStream(ds_streamId)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

	case "createtopic":
		createTopicCmd.Parse(os.Args[2:])
		ms := CreateMessageStream()
		if ct_name == "" {
			fmt.Println("Error: Name flag is required.")
			createTopicCmd.PrintDefaults()
			os.Exit(1)
		}

		err := ms.CreateTopic(ct_streamId, iggy.TopicRequest{
			TopicId:         ct_topicId,
			Name:            ct_name,
			PartitionsCount: ct_partitionsCount,
		})
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	case "gettopic":
		getTopicCmd.Parse(os.Args[2:])
		ms := CreateMessageStream()

		if gt_topicId == -1 {
			topics, err := ms.GetTopics(gt_streamId)
			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}
			jsonData, err := json.MarshalIndent(&topics, "", "  ")
			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}
			fmt.Println(string(jsonData))
			return
		}
		topic, err := ms.GetTopicById(gt_streamId, gt_topicId)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		jsonData, err := json.MarshalIndent(&topic, "", "  ")
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println(string(jsonData))

	case "deletetopic":
		deleteTopicCmd.Parse(os.Args[2:])
		ms := CreateMessageStream()
		if dt_streamId == -1 {
			fmt.Println("Error: Stream Id is required.")
			deleteStreamCmd.PrintDefaults()
			os.Exit(1)
		}

		if dt_topicId == -1 {
			fmt.Println("Error: Topic Id is required.")
			deleteStreamCmd.PrintDefaults()
			os.Exit(1)
		}

		err := ms.DeleteTopic(dt_streamId, dt_topicId)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

	case "help":
		getUsage()
	default:
		fmt.Println("Error: Unknown command.")
		getUsage()
		os.Exit(1)
	}
}

func CreateMessageStream() iggy.IMessageStream {
	factory := &iggy.MessageStreamFactory{}
	config := iggy.MessageStreamConfiguration{
		BaseAddress: url + ":" + port,
		Protocol:    iggy.Tcp,
	}

	ms, err := factory.CreateMessageStream(config)
	if err != nil {
		panic(err)
	}
	return ms
}
