#!/usr/bin/env bash
#title          :get_data:.sh
#description    :This script downloads the MOCO validation dataset 
#author         :Taha Azzaoui <tazzaoui@cs.uml.edu>
#version        :1    
#usage          :./get_data.sh
#================================================================

mkdir -p data 
cd data 

echo -e "Downloading Images...\n"
wget http://images.cocodataset.org/zips/val2014.zip
unzip val2014.zip
mv val2014 img
