package main

import (
	"log"
	"os"

	webview "github.com/webview/webview_go"
)

func main() {
	ArgParse(Run) //func ArgParse(f func(*Args) error)
}

func Run(args *Args) error {
	for i, a := range os.Args {
		switch i {
		case 0:
			continue
		default:
			a = a
			log.Printf("args: %v", a)
		}
	}

	if len(args.Url) == 0 {
		ShowHelp("Error: 引数にURLを与えてください。")
	}

	{
		// ウィンドウの幅と高さは最大最小値制限がある
		// らしいが、最大値が2になっている。
		log.Printf("args x, y: %v, %v", args.WinWidth, args.WinHeight)
		log.Printf("Min, Max: %v, %v", webview.HintMin, webview.HintMax)
		if args.WinWidth < webview.HintMin {
			args.WinWidth = webview.HintMin
			log.Printf("args x, y: %v, %v", args.WinWidth, args.WinHeight)
		}
		if args.WinHeight < webview.HintMin {
			args.WinHeight = webview.HintMin
			log.Printf("args x, y: %v, %v", args.WinWidth, args.WinHeight)
		}
		if false {
			if args.WinWidth > webview.HintMax {
				args.WinWidth = webview.HintMax
				log.Printf("args x, y: %v, %v", args.WinWidth, args.WinHeight)
			}
			if args.WinHeight > webview.HintMax {
				args.WinHeight = webview.HintMax
				log.Printf("args x, y: %v, %v", args.WinWidth, args.WinHeight)
			}
		}
	}

	//w := webview.New(false)
	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle(args.WinTitle)
	w.SetSize(args.WinWidth, args.WinHeight, webview.HintNone)
	//w.SetHtml("Thanks for using webview!")
	w.Navigate(args.Url)
	w.Run()
	return nil
}
