mkdir -p data 
cd data 

echo -e "Downloading Images...\n"
wget http://images.cocodataset.org/zips/val2014.zip
unzip val2014.zip
mv val2014 img
