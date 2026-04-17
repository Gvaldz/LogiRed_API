package core

import (
    "github.com/disintegration/imaging"
)

func FixImageOrientation(filePath string) error {
    img, err := imaging.Open(filePath, imaging.AutoOrientation(true))
    if err != nil {
        return err
    }
    return imaging.Save(img, filePath)
}