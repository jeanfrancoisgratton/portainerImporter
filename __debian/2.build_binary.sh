#!/usr/bin/env bash

PKGDIR=portainerImporter-`cat current_pkg_release`_amd64

mkdir -p ${PKGDIR}/opt/bin ${PKGDIR}/DEBIAN
mv control ${PKGDIR}/DEBIAN/
mv preinst ${PKGDIR}/DEBIAN/

echo "Building binary from source"
cd ../src
go build -o ../__debian/${PKGDIR}/opt/bin/portainerImporter .
strip ../__debian/${PKGDIR}/opt/bin/portainerImporter
chown 0:0 ../__debian/${PKGDIR}/opt/bin/portainerImporter

echo "Binary built. Now packaging..."
cd ../__debian/
dpkg-deb -b ${PKGDIR}
