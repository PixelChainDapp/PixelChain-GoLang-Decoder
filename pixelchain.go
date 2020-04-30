package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	store "github.com/pixelchaindapp/PixelChain-GoLang-Decoder/contracts"
)

const width int = 32
const height int = 32

// Configuration for the application
type Configuration struct {
	NetworkAddress   string
	ContractAddress  string
	ImageStoragePath string
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal(errors.New("Expected argument eth id"))
	}

	ethID, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		log.Fatalf("%q is not a valid eth id.\n", os.Args[1])
	}

	configuration, err := loadConfiguration()
	if err != nil {
		log.Fatal(err)
	}

	client, err := ethclient.Dial(configuration.NetworkAddress)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("Connected to %s", configuration.NetworkAddress))

	address := common.HexToAddress(configuration.ContractAddress)
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Contract loaded")

	art, err := instance.PixelChains(&bind.CallOpts{}, big.NewInt(ethID))
	if err != nil {
		log.Fatal(err)
	}

	colors := paletteToColorRGBA(art.Palette)
	generateImage(art.Name, configuration.ImageStoragePath, art.Data, colors)

	fmt.Println("Image generated")
}

func loadConfiguration() (Configuration, error) {
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal(err)
	}

	configuration := Configuration{}
	err = json.Unmarshal([]byte(file), &configuration)
	if err != nil {
		return configuration, err
	}

	return configuration, nil
}

func paletteToColorRGBA(palette []byte) []color.RGBA {
	var colors []color.RGBA

	for i := 0; i < len(palette); i += 3 {
		r := int(palette[i])
		g := int(palette[i+1])
		b := int(palette[i+2])

		color := color.RGBA{
			uint8(r),
			uint8(g),
			uint8(b),
			255,
		}

		colors = append(colors, color)
	}

	return colors
}

func generateImage(name string, path string, imgData []byte, colors []color.RGBA) {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	x := 0
	y := 0

	for i := 0; i < len(imgData); i++ {
		img.Set(x, y, colors[int(imgData[i])])
		x++
		if x > width-1 {
			x = 0
			y++
		}
	}

	f, err := os.Create(fmt.Sprintf("%s%s.png", path, name))
	if err != nil {
		log.Fatal(err)
	}

	png.Encode(f, img)
}
