# Maintainer: Seu Nome <alexandre@dev2learn.com>
pkgname=lingmoos-autostart
pkgver=1.0.0
pkgrel=1
pkgdesc="Lingmo Service"
arch=('x86_64')
url="https://seuwebsite.com"
license=('MIT')
depends=('xdotool' 'rofi' 'sxhkd')
source=("${pkgname}-${pkgver}.tar.gz")
md5sums=('SKIP')

build() {
    cd "$srcdir/${pkgname}-${pkgver}"
    go build -o lingmoos main.go
}

package() {
    cd "$srcdir/${pkgname}-${pkgver}"
    install -Dm755 lingmoos "$pkgdir/usr/local/bin/lingmoos"
    # Crie o diretório de configuração do usuário
    install -d "$pkgdir/etc/systemd/system"
    echo "${serviceContent}" > "$pkgdir/etc/systemd/system/lingmo.service"
}
