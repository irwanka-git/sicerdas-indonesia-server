#!/usr/bin/env bash
orientation=$3
size=$4
url=$1
path=$2
nomorseri=$5
jenis=$6
utama="utama"
lampiran="lampiran"
cover="cover"

echo "orientation: $orientation"
echo "size: $size"
echo "url: $url"
echo "path: $path"
echo "jenis: $jenis"

if [ "$jenis" = "$utama" ]; then
    echo "Cetak Utama."
    wkhtmltopdf --footer-spacing 3 -L 10 -R 10 -T 15 -B 15 --footer-left "Si Cerdas Indonesia" --footer-right $nomorseri --footer-font-size 8 --footer-center [page]/[topage] -O $orientation -s $size $url $path
fi

if [ "$jenis" = "$lampiran" ]; then
    echo "Cetak Lampiran."
    wkhtmltopdf --footer-spacing 3 -L 15 -R 15 -T 10 -B 10 --footer-left "Lampiran" --footer-right $nomorseri --footer-font-size 8 -O $orientation -s $size $url $path
fi

if [ "$jenis" = "$cover" ]; then
    echo "Cetak Cover."
    wkhtmltopdf --footer-spacing 0 -L 0 -R 0 -T 0 -B 0 -O $orientation -s $size $url $path
fi

