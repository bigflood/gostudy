// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"strings"
	"io"
	"io/ioutil"

	"github.com/spf13/cobra"
	"github.com/aws/aws-sdk-go/service/polly"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/tosone/minimp3"
	"github.com/hajimehoshi/oto"
)

// sayCmd represents the say command
var sayCmd = &cobra.Command{
	Use:   "say",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		voice, err := cmd.Flags().GetString("voice")
		if err != nil {
			return err
		}

		svc, err := NewPolly()
		if err != nil {
			return err
		}

		input := &polly.SynthesizeSpeechInput{
			//LexiconNames: []*string{
			//	aws.String("example"),
			//},
			OutputFormat: aws.String("mp3"),
			SampleRate:   aws.String("8000"),
			Text:         aws.String(strings.Join(args, " ")),
			TextType:     aws.String("text"),
			VoiceId:      aws.String(voice),
		}

		result, err := svc.SynthesizeSpeech(input)
		if err != nil {
			return err
		}

		defer result.AudioStream.Close()

		//f, err := os.Create("gosay-output.mp3")
		//if err != nil {
		//	return err
		//}
		//
		//defer f.Close()
		//
		//numBytes, err := io.Copy(f, result.AudioStream)
		//if err != nil {
		//	return err
		//}
		//
		//fmt.Printf("%d bytes written\n", numBytes)

		return playMP3(result.AudioStream)
	},
}

func playMP3(data io.ReadCloser) error {
	bytes, err := ioutil.ReadAll(data)
	if err != nil {
		return err
	}
	dec, s, err := minimp3.DecodeFull(bytes)
	if err != nil {
		return err
	}

	player, err := oto.NewPlayer(dec.SampleRate, dec.Channels, 2, 1024)
	if err != nil {
		return err
	}

	if _, err := player.Write(s); err != nil {
		return err
	}

	return nil
}

func init() {
	rootCmd.AddCommand(sayCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sayCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	sayCmd.Flags().StringP("voice", "v", "Seoyeon", "select voice id")
}
