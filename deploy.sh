#!/bin/bash

# Jalankan git pull dengan sudo
sudo git pull

# Jalankan docker-compose up dengan build
sudo docker-compose up -d --build

# Membersihkan gambar yang tidak digunakan
sudo docker image prune -f