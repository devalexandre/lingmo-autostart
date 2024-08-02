# Maintainer: Seu Nome <seuemail@dominio.com>
pkgname=lingmoos-autostart
pkgver=1.0.0
pkgrel=1
pkgdesc="Lingmo Service"
arch=('x86_64')
url="https://seuwebsite.com"
license=('MIT')
source=("lingmo.service")
md5sums=('SKIP')

pre_install() {
    echo "Stopping the lingmo service..." | tee -a /tmp/pre_install_log.txt
    sudo systemctl stop lingmo.service | tee -a /tmp/pre_install_log.txt
    sudo rm -f "/etc/systemd/system/lingmo.service" | tee -a /tmp/pre_install_log.txt
}
package() {
    if [ -f "$pkgdir/etc/systemd/system/lingmo.service" ]; then
        echo "Removing existing lingmo.service..."
        sudo rm -f "$pkgdir/etc/systemd/system/lingmo.service"
    fi
    install -Dm644 "$srcdir/lingmo.service" "$pkgdir/etc/systemd/system/lingmo.service"
}

post_install() {
    echo "Enabling and starting the lingmo service..." | tee -a /tmp/post_install_log.txt
    sudo systemctl daemon-reload | tee -a /tmp/post_install_log.txt
    sudo systemctl enable lingmo.service | tee -a /tmp/post_install_log.txt
    sudo systemctl start lingmo.service | tee -a /tmp/post_install_log.txt
}
