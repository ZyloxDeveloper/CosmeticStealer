package main

import (
	"encoding/json"
	"fmt"
	"image"
	"image/png"
	"os"
	"strings"

	"github.com/sandertv/gophertunnel/minecraft"
	"github.com/sandertv/gophertunnel/minecraft/auth"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
	"golang.org/x/oauth2"
)

var (
	RemoteAddress string
	TokenSource   oauth2.TokenSource

	AssetPath string
)

func main() {
	TokenSource = auth.RefreshTokenSource(parseToken())

	fmt.Print("\nPlease enter target address... Zeqa.net:19132\n> ")
	fmt.Scan(&RemoteAddress)

	AssetPath = "assets/" + strings.Replace(RemoteAddress, ":", "-", -1)

	os.Mkdir("assets", 0777)
	os.Mkdir(AssetPath, 0777)

	serverConn, err := minecraft.Dialer{
		TokenSource: TokenSource,
	}.Dial("raknet", RemoteAddress)
	if err != nil {
		panic(err)
	}

	if err := serverConn.DoSpawn(); err != nil {
		panic(err)
	}

	fmt.Printf("> %s is now online.\n", serverConn.IdentityData().DisplayName)

	for {
		pk, err := serverConn.ReadPacket()

		if pl, ok := pk.(*packet.PlayerList); ok {
			for _, player := range pl.Entries {
				if player.Username != serverConn.IdentityData().DisplayName && player.Username != "" {
					appendSkin(player.Username, player.Skin)

					fmt.Printf("> Parsing %s...\n", player.Username)
				}
			}
		}

		if err != nil {
			panic(err)
		}
	}
}

var DefaultSkinGeometry = `{"format_version":"1.12.0","minecraft:geometry":[{"bones":[{"name":"body","parent":"waist","pivot":[0,24,0]},{"name":"waist","pivot":[0,12,0]},{"cubes":[{"origin":[-5,8,3],"size":[10,16,1],"uv":[0,0]}],"name":"cape","parent":"body","pivot":[0,24,3],"rotation":[0,180,0]}],"description":{"identifier":"geometry.cape","texture_height":32,"texture_width":64}},{"bones":[{"name":"root","pivot":[0,0,0]},{"cubes":[{"origin":[-4,12,-2],"size":[8,12,4],"uv":[16,16]}],"name":"body","parent":"waist","pivot":[0,24,0]},{"name":"waist","parent":"root","pivot":[0,12,0]},{"cubes":[{"origin":[-4,24,-4],"size":[8,8,8],"uv":[0,0]}],"name":"head","parent":"body","pivot":[0,24,0]},{"name":"cape","parent":"body","pivot":[0,24,3]},{"cubes":[{"inflate":0.5,"origin":[-4,24,-4],"size":[8,8,8],"uv":[32,0]}],"name":"hat","parent":"head","pivot":[0,24,0]},{"cubes":[{"origin":[4,12,-2],"size":[4,12,4],"uv":[32,48]}],"name":"leftArm","parent":"body","pivot":[5,22,0]},{"cubes":[{"inflate":0.25,"origin":[4,12,-2],"size":[4,12,4],"uv":[48,48]}],"name":"leftSleeve","parent":"leftArm","pivot":[5,22,0]},{"name":"leftItem","parent":"leftArm","pivot":[6,15,1]},{"cubes":[{"origin":[-8,12,-2],"size":[4,12,4],"uv":[40,16]}],"name":"rightArm","parent":"body","pivot":[-5,22,0]},{"cubes":[{"inflate":0.25,"origin":[-8,12,-2],"size":[4,12,4],"uv":[40,32]}],"name":"rightSleeve","parent":"rightArm","pivot":[-5,22,0]},{"locators":{"lead_hold":[-6,15,1]},"name":"rightItem","parent":"rightArm","pivot":[-6,15,1]},{"cubes":[{"origin":[-0.1,0,-2],"size":[4,12,4],"uv":[16,48]}],"name":"leftLeg","parent":"root","pivot":[1.9,12,0]},{"cubes":[{"inflate":0.25,"origin":[-0.1,0,-2],"size":[4,12,4],"uv":[0,48]}],"name":"leftPants","parent":"leftLeg","pivot":[1.9,12,0]},{"cubes":[{"origin":[-3.9,0,-2],"size":[4,12,4],"uv":[0,16]}],"name":"rightLeg","parent":"root","pivot":[-1.9,12,0]},{"cubes":[{"inflate":0.25,"origin":[-3.9,0,-2],"size":[4,12,4],"uv":[0,32]}],"name":"rightPants","parent":"rightLeg","pivot":[-1.9,12,0]},{"cubes":[{"inflate":0.25,"origin":[-4,12,-2],"size":[8,12,4],"uv":[16,32]}],"name":"jacket","parent":"body","pivot":[0,24,0]}],"description":{"identifier":"geometry.humanoid.custom","texture_height":64,"texture_width":64,"visible_bounds_height":2,"visible_bounds_offset":[0,1,0],"visible_bounds_width":1}},{"bones":[{"name":"root","pivot":[0,0,0]},{"name":"waist","parent":"root","pivot":[0,12,0]},{"cubes":[{"origin":[-4,12,-2],"size":[8,12,4],"uv":[16,16]}],"name":"body","parent":"waist","pivot":[0,24,0]},{"cubes":[{"origin":[-4,24,-4],"size":[8,8,8],"uv":[0,0]}],"name":"head","parent":"body","pivot":[0,24,0]},{"cubes":[{"inflate":0.5,"origin":[-4,24,-4],"size":[8,8,8],"uv":[32,0]}],"name":"hat","parent":"head","pivot":[0,24,0]},{"cubes":[{"origin":[-3.9,0,-2],"size":[4,12,4],"uv":[0,16]}],"name":"rightLeg","parent":"root","pivot":[-1.9,12,0]},{"cubes":[{"inflate":0.25,"origin":[-3.9,0,-2],"size":[4,12,4],"uv":[0,32]}],"name":"rightPants","parent":"rightLeg","pivot":[-1.9,12,0]},{"cubes":[{"origin":[-0.1,0,-2],"size":[4,12,4],"uv":[16,48]}],"name":"leftLeg","parent":"root","pivot":[1.9,12,0]},{"cubes":[{"inflate":0.25,"origin":[-0.1,0,-2],"size":[4,12,4],"uv":[0,48]}],"name":"leftPants","parent":"leftLeg","pivot":[1.9,12,0]},{"cubes":[{"origin":[4,11.5,-2],"size":[3,12,4],"uv":[32,48]}],"name":"leftArm","parent":"body","pivot":[5,21.5,0]},{"cubes":[{"inflate":0.25,"origin":[4,11.5,-2],"size":[3,12,4],"uv":[48,48]}],"name":"leftSleeve","parent":"leftArm","pivot":[5,21.5,0]},{"name":"leftItem","parent":"leftArm","pivot":[6,14.5,1]},{"cubes":[{"origin":[-7,11.5,-2],"size":[3,12,4],"uv":[40,16]}],"name":"rightArm","parent":"body","pivot":[-5,21.5,0]},{"cubes":[{"inflate":0.25,"origin":[-7,11.5,-2],"size":[3,12,4],"uv":[40,32]}],"name":"rightSleeve","parent":"rightArm","pivot":[-5,21.5,0]},{"locators":{"lead_hold":[-6,14.5,1]},"name":"rightItem","parent":"rightArm","pivot":[-6,14.5,1]},{"cubes":[{"inflate":0.25,"origin":[-4,12,-2],"size":[8,12,4],"uv":[16,32]}],"name":"jacket","parent":"body","pivot":[0,24,0]},{"name":"cape","parent":"body","pivot":[0,24,-3]}],"description":{"identifier":"geometry.humanoid.customSlim","texture_height":64,"texture_width":64,"visible_bounds_height":2,"visible_bounds_offset":[0,1,0],"visible_bounds_width":1}}]}`

func appendSkin(n string, s protocol.Skin) {
	sid := s.SkinID
	if string(s.SkinGeometry) != DefaultSkinGeometry && string(s.SkinGeometry) != "null" && len(s.SkinGeometry) != 0 {
		os.Mkdir(AssetPath+"/"+n, 0777)
		f, err := os.Create(fmt.Sprintf("%s/%s/%s_skin.png", AssetPath, n, sid))
		if err != nil {
			fmt.Printf("> Error logging %s's skin...\n", n)
			return
		}

		writeSkinPng(int(s.SkinImageWidth), int(s.SkinImageHeight), s.SkinData, f)

		err = os.WriteFile(fmt.Sprintf("%s/%s/%s_geometry.json", AssetPath, n, sid), s.SkinGeometry, 0644)
		if err != nil {
			fmt.Printf("> Error logging %s's skin data...\n", n)
			return
		}

		fmt.Printf("> Logged %s's skin.\n", n)
	}

	if string(s.CapeData) != "null" && len(s.CapeData) != 0 {
		os.Mkdir(AssetPath+"/"+n, 0777)
		f, err := os.Create(fmt.Sprintf("%s/%s/%s_cape.png", AssetPath, n, sid))
		if err != nil {
			fmt.Printf("> Error logging %s's cape...\n", n)
			return
		}

		writeSkinPng(int(s.CapeImageWidth), int(s.CapeImageHeight), s.CapeData, f)

		fmt.Printf("> Logged %s's cape.\n", n)
	}
}

func writeSkinPng(w, h int, pix []byte, f *os.File) {
	if w == 0 || h == 0 {
		return
	}

	img := image.NewRGBA(image.Rect(0, 0, int(w), int(h)))
	img.Pix = pix

	err := png.Encode(f, img)
	if err != nil {
		panic(err)
	}

	f.Close()
}

func parseToken() *oauth2.Token {
	var t *oauth2.Token
	b, _ := os.ReadFile("token.tok")
	json.Unmarshal(b, &t)

	if !t.Valid() {
		t, _ = auth.RequestLiveToken()
		b, _ = json.Marshal(t)
		os.WriteFile("token.tok", b, 0644)
	}

	return t
}
