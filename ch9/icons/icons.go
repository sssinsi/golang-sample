package icons

import (
	"image"
	"sync"
)

// var mu sync.Mutex //iconsを保護する
var mu sync.RWMutex //iconsを保護する
var icons map[string]image.Image

func loadIcons() {
	icons = map[string]image.Image{
		"spades.png":   loadIcon("spades.png"),
		"hearts.png":   loadIcon("hearts.png"),
		"diamonds.png": loadIcon("diamonds.png"),
		"clubs.png":    loadIcon("clubs.png"),
	}
}

func loadIcon(name string) image.Image { return nil }

//注意 並行的に安全ではない。遅延初期化を行っている
// func Icon(name string) image.Image {
// 	if icons == nil {
// 		loadIcons() //1回だけの初期化
// 	}
// 	return icons[name]
// }

//並行的に安全、iconsへの相互排他アクセスを強制する。
// func Icon(name string) image.Image {
// 	mu.Lock()
// 	defer mu.Unlock()
// 	if icons == nil {
// 		loadIcons()
// 	}
// 	return icons[name]
// }

//並行的に安全
// func Icon(name string) image.Image {
// 	mu.RLock()
// 	if icons != nil {
// 		icon := icons[name]
// 		mu.RUnlock()
// 		return icon
// 	}
// 	mu.RUnlock()

// 	//排他的なロックを獲得する
// 	mu.Lock()
// 	if icons == nil { //注意:再びnil検査なければならない
// 		loadIcons()
// 	}
// 	icon := icons[name]
// 	mu.Unlock()
// 	return icon
// }

var loadIconsOnce sync.Once

func Icon(name string) image.Image {
	loadIconsOnce.Do(loadIcons) //呼び出しごとにミュータックスをロックしてブーリアン変数を検査する。
	return icons[name]
}
