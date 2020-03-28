#!/usr/bin/env python
import os
import random

def pull_data():
    os.system("wget http://www.vision.caltech.edu/Image_Datasets/Caltech101/101_ObjectCategories.tar.gz")
    os.system("tar -xvf 101_ObjectCategories.tar.gz")
    os.system("rm -rf 101_ObjectCategories/BACKGROUND_Google")

def sample_data(k):
    os.system("rm -rf img && mkdir -p img")

    for category in os.listdir("101_ObjectCategories"):
        count = 0
        path = os.path.join("101_ObjectCategories", category)
        if os.path.isdir(path):
            images = [i for i in os.listdir(path) if i.endswith(".jpg")]
            sample = random.choices(images, k=k)

            for i, s in enumerate(sample):
                src = os.path.join(path, s)
                dst = os.path.join("img", "{}_{}.jpg".format(category, i))
                os.system("cp {} {}".format(src, dst))

if __name__ == "__main__":
    if not os.path.exists("101_ObjectCategories"):
        pull_data()
    else:
        print("Sampling from existing data...")
    sample_data(k=5)
