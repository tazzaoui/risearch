package config

// Path to reference bank
const IMG_DIR = "data/img"

// Output path of their extracted keypoints
const KP_DIR = "data/kp"

// Maximum number of images index
const MAX_IMAGES = 100

// Count of best matches found per each query descriptor. Used in knnMatch()
// See: https://docs.opencv.org/master/db/d39/classcv_1_1DescriptorMatcher.html#aa880f9353cdf185ccf3013e08210483a
const K = 4
