# risearch

Writing reverse image search to learn go

![SIFT Keypoints](screenshot.png)

## Deps
* [gocv](https://gocv.io/)

## Basic idea:
* Construct a matrix of SIFT descriptors of each reference image
* Compute SIFT descriptors of an input image
* Compare similarity across all reference images (kmeans?)
