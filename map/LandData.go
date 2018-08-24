package main

import (
	"fmt"
	"sort"
)

type Resource struct {
	gold, wood, water, fire, earth int
}
// 分别找到金木水火土五种元素中的最大值，并且把找出不重复的元素

var res = [2025]Resource{{5,1,89,4,1},{5,2,88,2,3},{5,92,3,91,2},{1,82,4,5,5},{1,2,99,1,99},{3,2,1,97,94},{96,2,2,93,5},{5,2,89,1,2},{1,93,5,4,92},{2,2,2,5,89},{3,85,3,4,2},{3,2,4,89,133},{3,2,92,2,133},{2,2,2,2,133},{4,1,97,96,133},{2,131,99,97,1},{5,1,83,5,5},{4,1,97,95,1},{1,2,3,97,96},{3,2,90,3,3},{4,2,96,3,94},{1,92,1,1,2},{5,80,4,5,4},{2,1,2,89,5},{4,2,3,87,2},{2,2,97,2,97},{98,2,2,3,95},{92,1,4,91,5},{2,2,98,95,2},{94,2,96,2,4},{1,2,94,1,4},{2,86,4,3,3},{1,2,4,90,4},{2,2,97,5,93},{3,88,2,1,3},{93,2,3,1,94},{5,2,133,3,2},{4,1,96,3,92},{2,2,3,85,5},{99,2,1,97,1},{92,1,5,95,4},{4,91,94,5,3},{5,2,2,4,87},{4,2,1,4,88},{93,1,95,5,4},{4,92,1,5,95},{96,93,3,2,4},{5,2,86,2,4},{87,2,4,2,3},{2,93,4,4,94},{2,2,2,2,92},{2,94,3,96,1},{3,2,4,1,88},{4,1,93,2,133},{2,2,93,3,135},{1,1,98,97,144},{91,2,5,1,150},{5,2,5,1,150},{4,133,1,93,150},{5,142,5,1,87},{2,148,2,5,144},{1,142,4,95,135},{95,133,2,94,2},{5,1,4,85,2},{95,2,3,2,97},{2,1,97,2,97},{2,2,96,4,96},{1,94,96,4,3},{87,2,5,3,1},{5,1,96,1,96},{97,1,4,1,96},{85,2,4,5,2},{3,1,90,1,4},{4,2,1,2,90},{4,2,96,95,2},{1,2,97,4,96},{1,2,2,89,3},{2,95,2,3,93},{95,131,5,5,91},{5,131,135,2,1},{89,2,144,2,3},{2,1,150,96,1},{3,1,144,98,98},{95,1,135,3,96},{2,84,3,4,4},{91,2,5,93,5},{93,2,2,92,4},{89,1,3,4,1},{87,1,5,4,1},{91,92,2,5,3},{93,2,2,5,93},{1,90,3,3,1},{4,2,5,85,2},{1,2,91,4,2},{3,80,5,5,4},{1,2,4,4,86},{1,2,4,4,135},{95,1,5,95,144},{96,1,97,5,150},{4,2,4,89,150},{94,1,4,94,170},{5,131,1,1,90},{3,133,95,5,193},{96,142,3,5,195},{5,168,5,4,186},{2,197,99,1,170},{5,168,2,88,144},{95,148,99,1,2},{5,142,96,95,2},{88,133,2,3,5},{4,2,4,3,87},{2,1,88,3,5},{2,92,5,4,91},{90,2,2,1,5},{4,86,2,3,3},{84,2,5,4,3},{3,1,1,87,3},{96,2,1,95,5},{1,86,4,3,4},{4,92,94,4,2},{5,2,1,3,87},{94,133,96,5,3},{93,142,4,96,3},{1,148,135,1,3},{4,148,144,1,91},{3,142,93,1,2},{1,133,199,1,98},{4,2,170,2,88},{2,2,144,1,95},{2,1,135,93,4},{5,2,5,89,93},{5,92,1,4,92},{97,2,97,4,1},{2,1,97,95,1},{94,2,5,91,5},{87,1,5,3,4},{89,90,5,5,5},{5,81,3,4,5},{2,2,2,95,98},{5,1,2,91,135},{1,1,3,93,144},{90,2,1,4,150},{3,1,92,3,170},{3,2,99,99,199},{94,133,5,1,194},{4,142,2,3,188},{3,1,1,1,392},{5,1,3,3,384},{2,2,2,2,390},{5,1,4,1,386},{1,390,4,1,1},{5,194,96,3,150},{1,196,98,2,133},{93,168,5,94,4},{3,142,97,3,95},{1,2,3,88,5},{88,2,4,1,1},{3,93,4,4,92},{5,1,5,4,85},{94,2,1,94,3},{2,1,2,4,88},{89,2,4,3,1},{5,2,94,5,92},{4,1,95,96,3},{5,2,89,2,3},{97,133,1,94,5},{95,142,133,1,2},{4,168,133,1,93},{5,191,144,2,133},{4,191,170,2,133},{2,168,195,2,133},{1,2,388,4,4},{4,133,194,1,133},{96,1,170,1,1},{98,2,144,96,5},{94,2,97,1,3},{4,2,5,82,4},{5,90,93,5,4},{5,84,2,4,2},{5,92,1,4,91},{94,1,2,94,2},{89,1,4,1,1},{1,2,1,92,1},{3,2,95,2,135},{1,2,100,97,144},{88,1,3,5,170},{100,1,1,100,200},{3,2,99,99,196},{2,1,4,4,387},{4,2,4,4,387},{5,1,4,2,386},{5,2,4,3,382},{2,2,2,4,391},{3,2,5,5,387},{2,381,4,5,5},{1,389,3,3,2},{3,386,3,2,4},{3,385,1,5,3},{2,193,95,2,4},{97,148,2,99,1},{3,131,98,1,94},{1,87,1,4,5},{4,93,4,2,92},{1,2,4,90,3},{96,1,95,4,3},{3,2,86,5,4},{96,1,5,92,3},{3,1,1,3,90},{95,2,3,92,1},{2,2,135,2,92},{1,142,144,4,2},{5,168,150,4,135},{1,193,150,3,144},{2,389,1,4,1},{3,386,4,1,4},{1,2,394,1,4},{2,2,392,3,3},{1,2,390,4,3},{4,2,190,4,144},{1,1,150,95,135},{1,2,133,5,4},{89,2,2,2,3},{84,1,3,4,4},{90,2,5,92,5},{2,2,3,90,1},{3,2,5,2,89},{5,2,5,83,3},{90,2,5,4,135},{86,2,3,5,144},{4,1,4,87,170},{5,1,97,97,195},{2,2,2,1,394},{1,2,1,3,392},{2,2,5,1,387},{1,1,3,2,393},{1,1,2,3,392},{4,383,2,4,5},{5,380,3,5,4},{4,387,1,3,3},{4,389,2,1,1},{3,385,4,5,1},{4,193,2,5,94},{99,197,1,100,1},{1,168,98,98,1},{5,142,94,5,93},{1,2,4,92,3},{3,1,87,4,5},{2,2,86,4,5},{3,2,3,83,5},{4,2,90,3,2},{85,2,3,4,4},{2,2,1,87,4},{4,2,95,96,2},{5,2,2,83,4},{1,131,144,5,94},{5,148,170,5,88},{5,194,196,1,144},{4,387,5,1,1},{3,385,3,2,5},{1,386,4,3,4},{5,387,2,2,1},{5,1,388,3,2},{1,2,390,4,2},{4,131,193,3,170},{2,2,150,5,144},{2,2,133,97,96},{3,2,94,2,1},{5,2,1,95,97},{88,2,5,2,5},{85,2,2,4,5},{98,2,1,95,3},{3,2,1,92,1},{4,2,5,90,144},{1,2,5,92,170},{1,1,3,98,194},{3,2,2,2,387},{5,1,5,4,381},{1,2,2,2,394},{4,148,4,4,85},{5,191,4,5,192},{5,382,1,4,5},{5,385,2,4,1},{2,196,99,1,150},{2,195,5,93,150},{4,186,3,5,144},{2,195,2,98,135},{4,168,5,83,5},{92,148,4,5,92},{1,142,94,3,2},{90,133,2,2,3},{2,2,3,5,87},{5,2,1,85,2},{5,1,95,5,94},{4,2,88,4,4},{5,2,87,3,2},{86,2,4,3,3},{93,2,3,2,95},{96,1,1,1,99},{98,1,133,3,97},{99,2,150,1,95},{2,142,192,5,133},{4,1,389,2,4},{5,2,387,1,4},{1,2,5,3,387},{2,1,5,5,387},{5,2,5,2,388},{3,2,1,4,388},{1,2,1,5,387},{4,1,96,3,193},{5,1,144,3,150},{5,2,90,4,133},{3,1,4,3,86},{1,2,4,5,87},{1,1,1,4,92},{96,2,1,93,4},{94,1,3,95,3},{1,2,95,3,133},{88,1,4,5,150},{3,2,99,96,198},{4,2,3,4,387},{5,2,2,3,386},{97,2,3,97,197},{97,2,1,97,196},{1,142,5,90,170},{89,168,4,4,150},{93,194,2,5,150},{92,194,4,5,144},{5,168,90,4,135},{91,148,5,1,133},{2,148,98,3,96},{1,148,98,1,98},{3,142,94,1,2},{5,133,88,2,4},{94,2,4,91,4},{95,2,4,94,3},{92,1,4,3,92},{5,1,84,5,4},{2,2,1,4,90},{84,1,5,2,5},{1,2,1,98,95},{94,2,1,94,3},{3,2,89,2,3},{90,2,3,4,1},{5,2,133,5,4},{89,2,150,3,4},{1,133,92,3,4},{2,1,392,2,2},{1,1,395,1,1},{3,2,388,3,3},{4,1,2,2,389},{3,2,3,5,383},{1,2,4,2,390},{3,2,4,5,386},{96,1,144,1,199},{1,1,93,5,150},{90,2,5,1,133},{88,2,2,1,4},{2,2,3,5,87},{2,1,93,2,1},{3,2,91,1,4},{4,1,1,3,87},{3,2,4,88,133},{2,2,93,2,150},{98,2,1,97,197},{5,2,4,1,388},{4,2,5,1,186},{4,2,90,4,170},{96,2,4,92,150},{92,133,4,3,144},{92,142,1,5,135},{97,148,1,2,133},{93,148,4,5,96},{5,142,5,2,85},{97,133,3,97,1},{4,131,96,4,96},{1,131,97,4,97},{2,2,3,2,88},{3,2,4,88,3},{1,2,1,2,92},{4,2,2,86,5},{3,1,1,2,92},{89,1,5,1,3},{2,1,2,5,87},{95,1,2,96,1},{1,2,4,1,91},{2,2,1,95,98},{3,2,2,5,86},{85,2,2,5,4},{2,2,1,88,2},{3,2,144,95,99},{2,2,170,3,2},{5,1,91,2,135},{4,2,384,4,5},{4,2,389,2,5},{2,2,393,2,3},{2,2,5,4,389},{3,1,5,4,387},{95,2,144,4,195},{93,2,2,5,170},{4,1,93,2,144},{92,2,95,3,5},{1,1,4,91,1},{3,2,5,85,4},{1,2,5,92,95},{2,2,3,92,96},{5,2,3,5,85},{5,2,1,87,4},{2,1,4,4,144},{5,1,1,93,170},{99,1,98,4,196},{5,2,3,91,170},{5,1,3,90,144},{1,2,96,3,135},{2,1,1,3,90},{2,2,90,5,2},{3,131,97,1,95},{5,131,3,95,94},{5,2,94,92,5},{2,1,4,92,95},{3,1,2,90,2},{5,2,1,88,2},{5,2,2,95,94},{5,2,1,3,89},{3,2,1,5,91},{5,2,85,4,5},{89,2,4,3,2},{4,1,2,87,1},{88,2,5,1,4},{4,2,5,85,1},{5,2,5,2,82},{90,2,2,1,2},{95,1,5,91,4},{92,1,94,5,4},{94,2,5,93,5},{97,1,135,93,4},{98,2,144,5,98},{92,2,170,4,2},{2,1,198,94,135},{2,2,388,5,4},{1,2,390,5,3},{4,1,388,5,1},{5,2,196,3,195},{1,2,150,2,170},{94,2,133,2,144},{5,2,90,5,135},{94,1,3,3,94},{94,2,96,2,2},{83,1,5,4,4},{4,1,4,84,3},{1,1,92,4,1},{90,2,2,3,3},{84,1,5,5,3},{3,1,94,1,135},{2,2,97,94,144},{5,2,93,2,150},{88,1,3,5,144},{98,2,2,95,135},{1,2,94,1,1},{3,1,5,5,85},{3,2,3,3,86},{97,2,97,3,2},{2,1,4,92,96},{85,2,4,3,5},{2,1,133,91,4},{4,2,133,98,97},{85,2,5,5,2},{5,1,88,3,3},{5,2,2,85,5},{97,1,3,94,1},{88,1,5,2,2},{3,2,96,92,4},{96,1,3,4,96},{5,2,4,82,3},{3,2,96,1,94},{3,1,88,5,2},{1,2,4,95,94},{99,1,98,1,1},{1,2,1,95,96},{5,1,87,3,4},{90,2,1,5,1},{3,2,135,95,97},{4,2,144,3,90},{4,1,170,89,4},{5,1,196,1,97},{3,2,94,3,144},{2,2,199,95,150},{95,2,170,1,150},{92,2,144,3,144},{2,2,95,1,135},{2,2,3,5,89},{4,2,3,96,95},{85,1,5,4,1},{95,2,5,93,3},{4,1,90,1,3},{3,2,90,1,4},{96,2,3,93,2},{90,1,3,2,3},{3,2,94,2,1},{93,2,1,2,1},{2,1,5,92,133},{93,2,2,94,2},{82,2,4,5,5},{3,2,97,93,2},{2,2,3,86,3},{4,2,85,3,5},{90,2,3,2,3},{94,1,135,4,95},{2,2,144,1,94},{1,2,150,90,3},{88,1,150,2,5},{4,1,144,2,4},{5,2,135,92,93},{95,2,5,95,5},{3,1,2,90,2},{4,2,87,1,5},{2,2,97,93,1},{84,2,5,4,4},{3,2,2,95,96},{1,2,95,3,1},{4,1,96,95,4},{4,1,5,83,2},{88,2,3,1,3},{1,2,98,97,2},{5,2,2,88,3},{94,2,4,2,94},{4,1,2,89,1},{1,2,135,4,90},{94,2,144,2,2},{3,2,150,4,91},{5,2,150,98,95},{96,2,150,1,133},{5,2,144,3,133},{92,1,135,2,3},{5,1,2,92,94},{95,2,99,1,1},{83,2,4,5,2},{3,2,3,1,90},{1,2,98,97,1},{1,2,98,96,3},{2,2,5,3,88},{2,2,3,88,4},{99,2,99,1,1},{3,2,2,90,4},{2,2,91,1,5},{4,2,3,85,3},{1,2,5,93,93},{1,2,2,2,94},{1,2,94,4,1},{95,2,1,5,94},{91,2,5,1,1},{1,1,135,5,89},{92,1,144,1,5},{4,2,170,1,93},{5,2,196,96,1},{1,2,199,97,99},{3,2,170,5,87},{91,2,144,2,5},{93,2,1,92,5},{5,2,89,2,3},{89,2,3,5,1},{5,2,89,1,5},{3,1,98,93,1},{5,2,96,2,96},{3,2,5,2,90},{5,1,1,88,2},{1,2,4,5,88},{88,2,2,3,3},{1,2,1,94,95},{1,2,94,3,2},{3,2,3,82,5},{2,2,4,87,5},{2,2,4,88,3},{92,2,5,91,3},{96,2,133,1,97},{90,2,133,1,5},{5,2,133,95,95},{5,2,3,84,3},{85,2,2,5,4},{96,2,1,97,2},{87,2,2,1,5},{1,1,5,90,1},{92,2,3,96,3},{2,2,2,93,1},{4,1,88,3,3},{5,1,5,83,1},{2,2,3,87,3},{5,2,3,90,1},{3,2,5,87,3},{5,2,1,4,86},{5,2,93,5,90},{3,2,4,3,88},{90,1,4,4,1},{4,2,87,3,5},{5,2,86,3,5},{2,2,88,3,4},{5,2,144,85,5},{97,1,170,95,2},{5,2,196,94,1},{5,2,386,5,2},{2,2,393,1,2},{2,1,195,5,94},{95,1,150,95,2},{5,1,133,4,89},{2,1,4,85,5},{5,2,5,1,89},{1,2,92,3,2},{2,2,5,1,88},{2,1,88,3,5},{93,2,96,3,4},{2,1,1,93,2},{98,2,3,95,2},{97,2,97,2,2},{5,2,89,2,2},{90,1,1,4,4},{94,2,95,2,5},{95,2,3,94,4},{91,2,3,1,5},{4,1,91,1,3},{95,2,4,3,94},{1,1,5,90,2},{5,1,1,87,4},{2,2,3,3,89},{1,2,3,93,1},{2,1,97,93,1},{5,1,1,84,4},{5,2,89,2,2},{96,2,3,2,94},{4,2,94,90,4},{2,2,3,97,95},{89,2,3,5,1},{87,1,3,3,3},{85,2,5,2,3},{90,2,2,2,3},{1,1,4,96,97},{86,2,4,4,5},{94,2,97,5,2},{1,2,3,85,4},{1,2,5,89,5},{96,2,1,96,4},{90,1,133,2,3},{2,1,150,99,95},{2,2,198,93,2},{5,2,389,2,4},{2,2,394,1,1},{2,1,387,4,5},{3,1,198,1,95},{4,2,150,4,89},{1,2,133,92,1},{97,2,96,4,1},{4,2,89,5,2},{85,1,3,4,4},{4,1,3,4,87},{5,1,97,96,1},{92,2,3,2,1},{1,2,3,91,3},{3,1,97,2,95},{3,2,88,4,2},{97,1,3,96,2},{97,2,1,98,1},{94,1,4,1,95},{1,2,97,4,94},{5,2,86,5,4},{1,2,91,4,1},{91,2,4,95,5},{2,1,1,87,4},{1,1,90,5,3},{95,2,1,2,98},{1,2,5,89,4},{2,2,2,87,2},{3,1,96,4,93},{4,1,94,5,94},{2,1,92,2,2},{2,2,96,93,5},{97,2,4,97,3},{88,1,5,1,2},{2,2,90,5,2},{2,2,3,88,3},{2,1,4,4,85},{97,2,97,3,1},{4,2,2,94,94},{99,2,2,97,1},{3,1,2,3,89},{95,2,97,1,4},{89,2,2,4,5},{93,1,133,2,1},{97,2,150,3,96},{2,2,195,5,94},{3,2,393,2,2},{1,2,387,5,5},{1,1,390,5,2},{1,1,199,1,95},{92,2,150,3,3},{96,2,133,2,98},{91,2,3,1,3},{83,1,4,5,4},{2,2,4,5,89},{3,2,2,3,90},{95,1,2,5,92},{96,1,98,2,2},{89,2,3,1,4},{2,2,90,1,4},{5,1,4,88,2},{97,2,5,94,2},{90,2,4,1,3},{93,2,5,96,1},{86,2,5,2,1},{3,2,96,97,3},{5,2,87,4,2},{92,2,3,1,1},{3,2,94,90,5},{95,1,4,1,97},{5,2,1,86,4},{92,1,3,94,5},{133,2,95,1,4},{5,2,1,84,5},{83,2,4,4,4},{85,2,1,5,5},{4,1,4,84,4},{2,2,2,90,2},{2,2,97,1,95},{1,2,96,2,1},{5,1,88,1,4},{84,2,5,5,5},{89,2,5,2,1},{86,2,4,5,2},{3,2,95,4,95},{98,2,1,4,98},{94,1,4,5,94},{94,2,1,4,94},{5,1,5,85,3},{2,2,144,2,95},{88,2,170,4,5},{97,1,198,2,97},{2,2,394,2,2},{95,1,197,4,2},{4,2,170,87,5},{4,1,144,4,87},{3,1,3,86,3},{96,1,99,1,1},{94,1,5,90,4},{84,2,3,5,4},{5,2,2,87,2},{2,2,96,1,1},{5,2,87,3,5},{5,2,86,5,1},{97,2,3,2,96},{1,2,1,89,5},{2,2,98,94,1},{98,1,2,2,96},{4,1,96,4,96},{1,1,89,5,4},{93,2,3,1,3},{82,1,5,5,3},{4,2,3,85,3},{3,2,4,87,2},{95,2,97,2,3},{135,2,99,1,96},{144,2,4,3,5},{150,2,90,4,4},{144,2,92,3,3},{135,2,5,88,1},{5,2,2,89,1},{1,2,4,88,1},{4,2,90,1,2},{1,1,97,97,5},{95,1,4,93,2},{95,1,2,95,1},{2,2,3,5,88},{93,2,5,91,5},{1,1,98,95,1},{96,1,3,91,4},{4,2,4,87,4},{4,2,3,2,85},{95,1,96,3,5},{2,2,90,1,4},{2,2,135,90,3},{3,2,144,5,90},{3,2,170,96,98},{4,2,197,93,3},{90,2,170,4,2},{1,1,144,89,4},{1,2,135,4,94},{97,2,96,2,4},{89,2,5,1,1},{1,2,99,2,95},{5,2,88,3,3},{93,2,2,93,3},{94,2,4,1,95},{1,2,98,96,2},{1,1,1,88,5},{3,1,2,2,91},{89,1,4,2,3},{3,2,89,4,4},{83,1,5,2,5},{1,1,91,3,4},{1,1,92,4,1},{3,1,96,5,95},{95,2,1,92,3},{3,1,1,93,1},{4,2,87,5,1},{135,1,5,5,1},{144,1,3,88,4},{170,1,94,3,1},{192,2,4,5,94},{170,2,3,4,89},{144,2,5,93,97},{135,131,1,5,90},{4,1,1,90,2},{1,1,97,3,96},{3,2,2,5,87},{1,2,4,3,92},{97,2,4,2,96},{1,1,1,90,3},{2,2,95,5,94},{2,1,2,86,5},{89,2,3,5,3},{85,1,4,4,2},{3,2,2,2,89},{2,2,93,1,2},{5,1,3,2,87},{84,2,1,4,5},{94,2,135,2,3},{1,2,144,99,97},{5,2,150,5,90},{3,2,144,93,4},{94,1,135,1,3},{2,1,88,4,5},{4,1,4,91,96},{94,195,97,3,5},{2,95,98,194,4},{198,96,99,3,1},{1,97,199,3,97},{94,196,2,5,96},{94,95,197,5,1},{5,95,98,3,196},{3,95,3,198,94},{97,95,197,5,3},{94,1,3,96,5},{5,1,89,2,2},{2,1,2,98,97},{1,1,90,3,4},{5,2,93,92,5},{2,1,95,5,96},{94,2,5,94,2},{2,2,98,3,94},{144,2,96,2,2},{170,1,99,1,97},{198,2,3,98,2},{390,2,1,4,3},{195,133,4,1,98},{170,142,96,1,3},{144,148,1,94,5},{135,142,3,4,1},{97,133,97,4,3},{94,2,2,5,97},{2,1,5,88,2},{1,2,5,86,4},{1,2,1,2,93},{96,2,4,1,93},{1,2,98,94,2},{90,2,5,5,90},{1,2,2,1,92},{92,1,5,95,2},{5,2,97,95,1},{5,1,3,3,84},{5,1,5,3,86},{4,2,90,3,1},{4,2,1,89,3},{91,1,133,4,3},{4,2,95,94,5},{4,1,88,5,1},{5,2,4,83,4},{5,2,2,1,89},{95,2,4,195,94},{195,2,98,5,96},{100,2,100,197,2},{2,1,98,194,95},{99,197,98,1,2},{1,97,100,198,1},{95,96,3,197,4},{96,96,5,3,196},{99,97,1,195,3},{85,2,4,3,4},{93,1,4,96,2},{5,2,95,1,93},{4,2,5,3,88},{96,2,96,3,3},{3,2,1,4,88},{3,1,4,3,86},{133,2,3,96,95},{150,2,2,1,95},{194,1,96,4,3},{384,2,3,3,5},{388,2,4,5,1},{384,2,5,5,3},{191,168,5,5,94},{170,194,4,5,95},{144,168,93,3,4},{2,142,4,3,89},{2,2,4,93,1},{98,2,97,1,4},{88,1,5,3,1},{5,2,3,89,1},{1,2,1,90,5},{3,1,3,2,91},{1,1,92,5,1},{4,2,86,3,4},{92,2,3,90,5},{88,1,4,2,5},{4,2,2,93,1},{2,2,95,1,2},{4,2,3,84,4},{2,1,88,4,4},{2,2,3,90,3},{2,2,5,88,1},{3,2,5,3,89},{2,2,96,5,92},{4,1,89,3,2},{1,97,3,198,95},{95,1,3,196,99},{198,95,97,5,1},{94,195,5,1,95},{2,196,99,3,97},{94,94,5,197,4},{95,95,3,198,5},{95,97,1,196,3},{5,1,97,197,98},{96,2,98,1,3},{4,2,1,1,93},{1,1,1,5,90},{1,2,98,93,3},{86,1,1,3,5},{93,1,5,90,4},{93,1,3,93,2},{133,1,2,2,91},{150,2,4,87,4},{199,1,99,97,1},{384,1,3,4,5},{387,2,1,5,4},{393,2,2,1,2},{388,2,5,2,2},{4,384,4,2,3},{150,197,1,96,2},{133,148,2,2,5},{97,131,4,93,1},{5,2,95,2,94},{94,2,5,97,2},{92,2,1,96,5},{1,2,91,3,5},{86,2,1,5,3},{1,1,98,2,96},{5,2,1,5,86},{4,2,3,86,4},{1,2,3,92,1},{4,1,89,4,1},{3,2,88,5,3},{2,2,91,3,2},{4,2,4,86,3},{5,1,88,1,5},{3,2,1,95,97},{4,1,95,4,96},{2,2,4,4,88},{2,1,3,92,96},{4,94,96,196,5},{99,96,1,198,3},{98,2,99,1,197},{3,95,97,195,5},{93,195,96,4,5},{2,2,99,194,97},{2,97,99,195,3},{98,196,4,3,94},{97,1,97,198,5},{4,2,3,86,1},{91,1,2,1,3},{90,2,2,1,5},{3,2,5,90,94},{88,2,4,2,5},{3,1,1,87,4},{84,1,4,4,3},{1,1,91,1,5},{144,2,97,2,95},{170,2,4,88,4},{89,142,2,5,2},{390,1,2,2,4},{391,1,3,3,1},{3,389,1,4,1},{170,193,4,5,95},{144,168,2,4,2},{3,142,4,1,88},{88,1,2,1,4},{88,1,3,4,4},{92,1,3,2,1},{4,1,96,97,2},{5,2,5,82,1},{82,2,4,3,5},{133,1,3,4,2},{2,1,95,1,1},{3,2,4,5,86},{88,1,5,3,1},{4,2,2,4,87},{1,2,1,1,93},{5,2,2,95,95},{96,1,2,93,5},{97,1,1,95,1},{93,2,4,1,1},{3,1,2,90,2},{92,1,3,1,2},{5,2,4,85,3},{97,2,99,194,4},{3,96,98,194,1},{97,2,1,199,99},{1,2,99,198,100},{98,197,4,1,99},{195,96,99,2,1},{196,2,98,5,97},{97,2,4,194,98},{2,1,99,198,96},{98,2,1,97,1},{1,2,2,91,2},{95,2,2,5,96},{2,1,89,5,2},{3,2,4,5,84},{90,2,1,5,1},{1,2,3,1,95},{88,1,5,1,2},{135,2,5,86,5},{144,133,5,4,88},{170,148,1,4,3},{195,195,5,95,1},{1,387,4,3,2},{2,387,3,2,3},{5,387,2,1,3},{135,194,4,4,96},{1,148,98,3,94},{5,131,96,92,3},{92,2,94,5,5},{5,2,1,89,2},{88,1,4,3,3},{135,2,2,4,94},{144,2,4,4,5},{150,1,92,3,4},{144,1,3,2,92},{135,1,96,95,5},{133,2,4,91,2},{1,2,96,4,94},{3,2,5,91,93},{95,1,5,93,3},{1,2,5,89,1},{2,1,1,94,94},{2,1,4,88,5},{96,2,96,3,3},{1,2,99,2,99},{5,2,95,94,4},{96,2,98,197,2},{97,94,5,195,5},{94,95,5,1,197},{95,2,5,198,96},{97,96,198,4,3},{98,2,99,196,5},{99,97,1,197,1},{96,2,98,198,3},{197,96,2,4,98},{5,2,95,94,2},{1,2,90,4,5},{4,1,4,84,2},{2,2,5,92,95},{4,1,1,89,1},{4,1,4,81,5},{4,2,88,1,5},{1,1,2,1,94},{5,2,87,5,1},{135,142,5,5,3},{144,168,94,1,5},{150,191,133,2,4},{2,386,1,5,3},{4,381,4,3,5},{1,392,1,1,2},{3,194,2,1,93},{94,148,94,5,5},{4,131,1,93,96},{95,2,3,91,3},{5,1,4,82,3},{135,1,2,99,96},{144,1,97,98,3},{170,1,99,1,96},{198,1,98,4,95},{170,1,4,89,1},{144,2,2,4,5},{150,1,96,5,93},{144,2,1,5,90},{135,1,5,2,3},{1,1,97,2,97},{3,2,89,2,5},{4,2,5,3,85},{96,2,2,95,1},{4,2,3,3,87},{5,2,4,83,5},{1,2,4,89,4},{94,95,5,195,3},{96,2,198,4,95},{1,97,3,199,96},{96,1,99,197,1},{4,96,4,197,94},{199,2,100,1,98},{96,1,198,4,97},{99,2,4,199,99},{98,96,198,3,1},{94,2,95,3,4},{2,2,97,92,4},{2,1,5,85,2},{90,2,1,3,3},{94,2,4,1,96},{2,2,4,89,3},{2,2,3,86,5},{4,2,1,93,94},{92,131,3,4,1},{4,148,135,1,91},{4,193,94,2,5},{3,386,1,2,5},{1,387,3,5,2},{5,383,4,1,4},{4,194,97,5,94},{97,168,96,1,5},{95,142,3,1,97},{96,2,1,1,2},{4,2,5,79,5},{3,2,5,1,87},{144,2,97,5,97},{170,2,1,90,4},{194,1,4,94,5},{387,1,5,2,3},{199,2,98,3,98},{170,2,4,2,94},{194,2,5,4,93},{170,2,2,91,5},{144,1,92,2,5},{93,1,2,92,4},{4,2,92,1,1},{88,1,2,2,5},{93,2,5,94,5},{98,2,1,98,2},{4,1,5,5,82},{4,2,3,95,97},{94,94,5,192,5},{1,97,100,197,1},{5,95,3,198,98},{3,96,4,198,98},{4,95,96,193,5},{1,2,99,199,100},{2,96,98,196,1},{2,1,99,199,96},{4,96,98,197,1},{4,2,1,5,86},{1,2,2,93,4},{91,2,3,5,1},{2,2,1,88,5},{89,1,4,1,2},{98,2,98,4,1},{85,1,5,2,5},{4,2,2,95,93},{90,131,135,1,5},{1,148,144,5,94},{2,186,170,4,5},{4,381,5,4,3},{4,386,4,2,1},{5,386,3,1,3},{1,194,5,1,97},{96,148,4,1,94},{1,133,1,4,90},{88,2,2,5,2},{3,2,94,5,95},{133,2,97,3,94},{150,2,97,2,1},{196,2,2,96,98},{389,1,2,1,4},{388,2,3,4,4},{386,2,4,4,2},{195,2,97,4,1},{390,1,4,1,1},{198,1,98,5,97},{150,1,2,4,92},{133,1,2,5,88},{2,1,97,3,94},{1,1,5,4,87},{4,2,96,96,1},{4,1,89,5,1},{4,1,88,4,3},{3,2,3,84,3},{198,96,97,5,2},{94,96,3,193,4},{1,96,99,194,3},{4,95,98,196,3},{1,96,99,199,3},{97,95,5,195,4},{95,96,3,3,196},{99,2,5,197,96},{95,96,197,4,2},{93,2,4,94,2},{4,1,5,83,3},{5,2,85,4,3},{88,2,3,2,4},{2,2,95,94,4},{92,1,1,4,1},{5,1,5,95,92},{5,2,3,5,84},{2,2,144,5,3},{1,142,90,4,4},{2,168,94,2,1},{5,1,388,2,4},{5,384,2,2,5},{94,195,150,3,2},{94,168,133,1,1},{88,142,3,5,3},{95,2,4,93,1},{3,2,91,1,4},{2,2,1,90,2},{1,2,88,5,5},{144,2,5,90,1},{170,1,5,5,86},{192,1,4,91,5},{388,2,5,4,3},{388,1,2,1,4},{389,1,5,3,1},{387,2,4,2,4},{195,2,97,4,1},{150,1,99,2,96},{133,2,93,2,4},{94,2,3,92,5},{98,2,1,96,2},{2,2,89,5,3},{4,2,94,4,92},{87,2,5,5,3},{3,2,91,2,4},{2,1,93,2,2},{96,2,4,92,3},{5,2,2,86,2},{94,1,5,1,94},{5,2,4,5,83},{89,1,1,3,3},{5,2,2,3,87},{4,2,3,85,3},{3,2,91,4,2},{3,1,3,86,5},{96,1,2,3,93},{5,2,87,3,5},{1,2,88,5,3},{86,2,3,5,3},{5,2,2,4,87},{2,1,96,5,95},{1,1,133,93,5},{93,133,150,3,2},{1,142,92,3,3},{4,2,390,4,1},{3,1,388,3,4},{97,195,198,4,3},{94,168,150,1,5},{1,142,133,96,2},{97,133,2,1,95},{93,2,95,4,5},{3,2,95,1,1},{4,1,88,3,3},{3,1,5,84,4},{135,1,2,4,89},{144,2,94,1,4},{170,1,90,5,3},{194,2,1,5,95},{390,2,1,1,3},{392,2,1,3,4},{391,2,4,1,1},{194,1,1,1,1},{150,2,98,97,5},{133,2,1,94,3},{5,1,86,2,5},{1,2,91,4,2},{96,1,98,3,1},{1,1,99,1,97},{5,2,97,1,96},{2,2,5,5,82},{96,2,4,91,3},{5,1,2,87,3},{5,1,3,83,4},{95,1,98,2,1},{93,2,3,5,92},{5,2,1,86,5},{2,1,3,4,89},{5,1,2,89,2},{3,1,85,5,5},{92,1,4,4,92},{4,2,3,85,5},{5,1,95,5,94},{93,2,5,95,4},{87,1,3,5,3},{96,1,97,2,1},{89,2,4,1,4},{91,2,133,4,1},{2,142,150,2,93},{3,168,197,3,96},{3,2,390,2,2},{4,383,5,3,3},{2,390,1,1,3},{2,188,144,3,5},{94,148,98,2,3},{3,131,4,1,89},{3,1,3,1,89},{1,2,4,89,2},{94,2,1,3,94},{5,1,4,83,4},{96,2,2,92,2},{135,1,98,2,99},{144,2,5,2,89},{170,2,1,3,93},{197,2,98,95,4},{391,1,3,2,1},{197,2,2,97,3},{170,2,4,88,2},{144,1,5,2,89},{88,2,1,5,2},{5,2,96,92,4},{3,1,3,91,1},{95,2,95,5,2},{3,2,3,87,5},{96,2,4,93,1},{94,2,4,5,94},{4,2,91,1,2},{3,1,96,94,2},{1,1,94,3,1},{3,2,4,1,86},{83,2,2,4,5},{2,2,98,3,95},{2,1,3,94,93},{3,1,94,1,1},{1,2,95,3,1},{2,2,4,2,91},{4,2,5,2,86},{1,2,5,1,90},{5,2,86,4,5},{2,1,1,2,91},{4,2,5,2,85},{2,2,4,88,4},{3,131,93,2,1},{91,148,144,5,1},{1,197,170,99,1},{2,386,1,4,5},{2,386,3,5,2},{1,385,5,4,3},{1,192,135,1,3},{5,148,94,90,5},{4,131,92,2,1},{1,1,90,3,5},{3,1,87,5,4},{84,1,5,3,3},{92,1,95,3,5},{2,2,95,4,95},{93,2,2,97,4},{135,2,91,3,5},{144,2,96,5,95},{170,1,5,4,88},{195,2,96,4,5},{170,2,5,96,96},{144,1,5,90,3},{135,2,95,1,2},{3,2,5,4,88},{89,2,2,5,2},{5,2,3,5,82},{3,1,3,88,2},{87,2,3,5,3},{1,1,97,96,2},{5,2,83,4,5},{4,1,88,3,4},{87,2,2,4,2},{3,1,88,2,5},{5,2,3,3,88},{4,2,2,4,84},{5,2,86,4,5},{2,2,1,3,91},{3,1,1,2,89},{92,1,4,91,3},{5,2,1,92,93},{1,1,98,2,95},{96,2,2,93,2},{5,2,88,3,2},{4,2,1,89,3},{90,2,3,2,4},{5,2,3,95,93},{1,2,97,95,4},{1,142,135,97,1},{93,168,144,3,2},{3,188,150,4,3},{4,383,2,4,5},{96,195,135,3,2},{2,168,96,5,95},{4,142,2,96,93},{4,1,94,91,5},{87,2,3,4,5},{97,1,5,97,1},{87,2,1,5,3},{5,2,86,1,5},{92,2,5,3,94},{1,2,1,91,2},{1,2,90,4,2},{135,1,99,99,2},{144,2,93,3,1},{150,1,2,2,93},{144,2,4,89,4},{135,1,1,1,97},{91,1,4,5,93},{5,2,4,95,95},{92,2,5,93,5},{94,2,1,93,5},{4,1,96,2,93},{89,1,1,3,4},{3,2,4,3,86},{3,2,1,88,3},{5,2,94,5,93},{2,1,98,1,97},{4,1,85,5,4},{4,2,4,87,2},{3,2,97,3,97},{95,2,98,2,2},{3,2,95,1,1},{87,2,5,2,4},{86,2,4,4,2},{82,1,5,5,4},{88,2,4,5,1},{1,1,3,92,2},{1,2,5,85,4},{2,2,1,94,2},{4,1,97,93,1},{4,2,4,82,5},{2,2,2,4,86},{97,133,3,98,2},{95,142,3,2,94},{87,168,133,5,4},{1,196,99,4,97},{3,168,98,98,1},{3,142,91,2,3},{3,133,4,89,1},{1,2,3,91,3},{2,2,90,4,2},{1,1,5,5,87},{96,2,2,95,1},{1,2,97,95,1},{3,2,93,1,1},{5,2,96,4,96},{1,2,97,3,96},{3,2,2,86,4},{1,1,97,4,94},{133,2,3,91,2},{98,2,2,95,1},{5,131,2,93,96},{1,131,5,90,1},{2,1,2,88,4},{93,2,97,2,3},{94,2,1,94,1},{98,2,2,1,95},{92,2,4,90,5},{3,2,4,89,2},{2,2,1,2,93},{93,2,5,93,1},{94,2,95,5,4},{4,2,89,4,2},{4,2,4,94,92},{97,2,2,97,3},{93,1,4,2,96},{2,2,94,1,1},{3,1,2,88,2},{3,2,4,3,88},{90,1,1,2,2},{3,1,2,93,95},{1,2,97,98,2},{5,2,4,96,93},{1,1,3,96,97},{84,2,5,2,4},{88,2,1,5,4},{4,2,95,3,95},{5,1,97,95,1},{93,133,96,4,4},{1,142,97,1,1},{2,148,2,3,90},{84,142,5,5,4},{92,133,3,94,5},{96,2,96,1,4},{2,1,93,2,2},{2,2,4,97,96},{3,2,2,89,2},{5,2,4,91,96},{1,2,4,86,5},{2,2,1,5,90},{5,1,2,88,3},{1,2,87,5,5},{5,2,5,83,2},{93,131,2,3,1},{2,133,4,3,88},{97,142,2,93,5},{5,148,4,93,96},{84,148,5,3,5},{86,142,2,5,5},{2,133,98,96,2},{5,1,4,84,1},{3,2,2,1,92},{5,2,84,3,5},{2,1,3,85,4},{94,2,2,92,4},{94,2,2,4,95},{1,2,2,91,1},{1,1,1,95,2},{98,1,97,1,2},{95,2,1,95,3},{1,2,93,2,1},{2,2,1,4,92},{2,2,3,83,5},{3,2,91,1,3},{3,2,99,133,95},{93,1,3,133,1},{5,2,1,3,89},{89,2,2,4,2},{96,1,3,96,5},{88,2,5,1,4},{2,2,88,3,4},{98,2,98,4,1},{4,1,5,83,3},{93,2,95,4,5},{88,2,2,3,5},{2,131,5,5,86},{2,2,2,90,4},{2,2,99,95,1},{5,1,4,91,95},{2,2,98,2,94},{95,2,1,96,4},{95,1,4,5,91},{5,2,89,1,5},{5,2,5,2,84},{2,2,97,95,2},{97,2,1,94,4},{2,133,4,88,5},{86,142,1,5,5},{2,148,99,98,1},{1,148,3,92,1},{94,168,3,91,5},{96,196,4,99,1},{1,196,1,2,97},{98,168,99,1,2},{3,148,91,1,4},{1,142,3,94,98},{95,133,2,5,95},{3,1,3,2,88},{4,2,4,2,86},{2,2,92,1,5},{90,2,4,1,1},{4,2,86,4,5},{5,2,3,5,82},{93,2,2,4,92},{97,2,2,95,3},{1,2,1,92,4},{1,2,2,88,4},{2,1,98,135,97},{4,2,90,144,4},{87,2,4,150,5},{4,2,91,150,5},{5,2,5,144,3},{100,2,1,135,98},{3,2,3,5,85},{5,2,2,3,87},{2,1,3,97,93},{97,2,1,94,4},{97,2,3,94,2},{95,2,2,97,3},{1,2,4,5,88},{86,2,4,4,2},{96,1,3,92,4},{2,131,1,98,99},{2,2,3,89,3},{5,2,89,1,4},{86,2,5,5,3},{2,2,1,3,92},{93,1,94,4,5},{86,2,5,5,1},{4,2,3,4,86},{3,133,4,87,2},{3,142,95,96,5},{3,168,4,85,5},{1,197,98,4,97},{4,196,97,93,3},{4,193,3,3,91},{4,390,1,1,2},{3,389,2,2,2},{1,195,2,96,1},{98,195,98,3,3},{98,168,3,95,1},{1,142,5,85,5},{2,1,96,94,3},{2,1,3,3,88},{4,1,1,94,93},{4,2,95,91,4},{96,2,1,94,4},{3,2,5,85,2},{5,1,3,87,1},{90,1,1,3,3},{2,2,5,87,2},{97,1,4,135,97},{5,2,91,144,4},{3,1,94,170,2},{97,1,2,198,99},{3,1,96,196,3},{2,2,1,170,1},{97,2,96,144,5},{2,2,1,135,1},{88,2,1,5,5},{2,2,89,3,4},{94,2,1,5,96},{97,2,3,93,3},{95,2,5,90,4},{97,1,98,1,1},{4,133,2,1,92},{1,142,4,89,3},{94,148,1,5,93},{1,142,96,97,5},{4,133,2,94,97},{4,2,3,92,1},{2,2,87,4,4},{86,1,4,5,1},{5,2,86,5,1},{4,1,88,3,4},{2,142,97,3,97},{3,168,96,3,96},{94,195,97,5,2},{4,383,5,1,5},{2,384,1,5,5},{4,383,5,5,1},{4,385,3,4,1},{1,385,4,5,2},{5,381,1,5,5},{3,388,2,3,2},{97,198,99,1,1},{5,148,90,3,1},{95,131,3,4,95},{1,2,91,5,1},{2,131,4,94,94},{96,131,2,96,1},{5,2,88,3,1},{5,1,4,4,84},{1,2,5,90,3},{95,1,1,96,5},{92,1,3,135,4},{5,2,5,144,86},{4,1,5,170,1},{93,2,3,195,4},{5,2,3,387,2},{2,1,1,389,4},{93,1,5,196,1},{2,1,99,170,99},{2,1,4,144,93},{2,2,1,90,3},{93,1,4,96,3},{95,2,1,91,5},{4,2,90,2,2},{98,2,3,98,2},{3,133,1,95,97},{93,142,96,5,4},{1,168,97,97,5},{98,196,3,5,98},{1,168,98,4,96},{96,142,4,94,1},{5,2,5,85,4},{94,2,1,1,4},{83,2,5,5,1},{88,2,4,3,1},{94,131,2,94,5},{3,148,97,93,3},{99,196,2,3,95},{5,382,5,2,4},{2,384,4,3,5},{1,389,4,2,2},{1,390,2,2,2},{2,384,5,3,4},{5,383,3,4,2},{1,387,3,2,5},{1,386,4,2,5},{5,196,98,97,2},{95,148,3,96,1},{2,133,5,2,89},{1,142,99,95,2},{5,148,3,93,95},{4,148,3,92,93},{97,142,2,2,96},{3,133,3,88,1},{97,2,2,96,3},{2,2,5,2,87},{91,1,3,144,2},{5,2,2,170,87},{97,1,5,194,94},{4,1,1,387,3},{2,2,5,381,5},{1,1,2,388,3},{3,2,1,388,5},{2,2,98,199,96},{94,2,2,150,1},{1,1,97,133,1},{5,1,3,92,93},{92,1,2,1,4},{90,1,3,1,1},{91,133,3,1,1},{5,142,5,83,5},{3,168,97,1,96},{4,196,2,93,94},{1,391,2,1,3},{98,196,97,4,3},{2,148,99,1,99},{92,131,3,96,5},{2,2,5,86,2},{4,2,95,93,5},{5,2,4,84,2},{3,2,1,89,3},{1,142,97,4,97},{1,168,97,95,3},{94,195,1,5,97},{4,382,4,4,3},{5,385,4,3,1},{5,385,5,2,1},{5,381,3,5,3},{3,381,4,5,4},{3,385,2,4,4},{5,383,4,5,1},{2,196,99,1,98},{95,148,5,94,4},{1,142,4,89,1},{93,168,4,96,4},{97,195,2,98,5},{2,193,96,4,2},{5,168,96,96,4},{4,148,87,4,5},{5,142,95,3,95},{2,133,93,133,4},{4,1,98,150,98},{96,2,99,196,1},{1,1,3,386,4},{1,2,5,390,2},{2,2,3,390,3},{5,2,4,387,4},{96,2,1,195,4},{4,2,4,170,1},{2,2,2,144,92},{1,1,94,1,2},{93,2,3,91,3},{3,2,96,93,3},{92,2,5,3,93},{94,142,4,3,96},{3,168,97,94,1},{96,195,2,4,96},{3,383,4,5,3},{4,387,4,1,2},{94,194,5,4,96},{4,148,96,96,5},{93,131,3,93,3},{85,2,4,5,2},{91,2,4,1,3},{2,2,96,92,5},{94,1,2,98,2},{4,133,97,93,1},{2,142,1,94,1},{4,168,5,96,95},{4,194,2,2,94},{93,194,4,3,1},{94,194,4,92,5},{2,389,3,1,2},{3,384,1,5,5},{2,388,3,2,3},{5,193,1,94,4},{1,168,98,94,3},{1,142,97,94,5},{5,148,2,96,94},{96,196,3,1,97},{3,389,1,3,2},{4,388,1,2,3},{2,195,4,97,1},{93,192,1,5,4},{3,168,98,94,1},{96,142,4,95,5},{98,2,3,144,99},{94,2,5,170,94},{5,2,94,191,3},{3,2,4,389,2},{5,2,4,383,1},{93,1,2,195,3},{2,2,4,170,92},{2,2,97,144,94},{3,2,91,135,5},{3,2,3,87,3},{84,1,3,4,5},{3,2,92,2,2},{93,133,3,4,95},{86,148,2,4,5},{96,196,3,95,5},{1,388,3,2,4},{2,386,3,2,4},{3,386,1,3,4},{97,196,2,93,5},{89,148,3,1,5},{5,131,95,96,4},{5,2,4,85,1},{1,1,2,93,2},{4,2,1,90,4},{3,1,94,1,1},{92,2,3,96,3},{94,133,4,5,95},{94,148,95,3,5},{1,195,5,95,96},{4,379,5,4,5},{96,196,5,94,3},{3,382,5,4,4},{3,391,1,2,1},{4,193,4,2,95},{4,168,5,86,5},{2,142,93,3,1},{94,133,97,5,2},{5,142,97,1,94},{4,168,98,1,95},{3,196,99,1,97},{4,381,5,5,3},{3,384,5,4,1},{5,382,2,5,3},{1,194,3,2,95},{2,148,98,95,2},{1,131,1,91,5},{2,2,3,144,95},{91,2,4,170,1},{4,2,97,197,96},{96,2,1,195,99},{94,1,1,170,4},{96,2,4,144,96},{92,2,5,135,1},{94,2,4,3,95},{93,2,3,1,3},{5,2,85,5,5},{89,2,1,4,2},{94,142,1,96,4},{4,168,5,4,87},{5,195,1,98,96},{2,385,5,3,2},{4,386,1,3,3},{94,196,98,3,5},{1,168,99,95,2},{2,142,93,2,3},{4,2,90,1,3},{3,2,1,98,98},{5,2,1,5,83},{93,2,1,1,2},{94,2,97,1,1},{96,2,4,5,93},{94,2,5,5,91},{3,142,98,94,1},{1,168,4,90,1},{3,196,4,94,95},{3,168,5,2,90},{96,196,98,1,3},{3,195,98,3,97},{94,168,5,1,93},{87,142,5,3,1},{1,133,97,98,3},{94,1,5,3,93},{2,133,96,94,5},{98,142,4,1,95},{85,168,4,5,5},{93,195,2,3,2},{1,193,2,5,95},{3,381,5,5,4},{97,196,1,96,2},{89,148,1,4,5},{93,131,5,5,93},{1,2,95,135,3},{2,1,3,144,90},{1,1,97,150,1},{1,2,5,150,92},{3,2,93,144,4},{1,2,93,135,3},{99,2,99,2,1},{5,2,2,82,4},{2,1,5,84,5},{83,1,5,3,5},{94,131,4,5,93},{97,148,3,1,98},{4,193,3,2,94},{5,382,5,3,3},{2,388,1,5,1},{3,191,5,91,4},{5,168,95,4,96},{5,142,5,89,1},{1,133,2,94,95},{1,2,2,1,94},{3,1,93,1,1},{5,1,90,2,2},{93,2,3,92,5},{2,2,4,86,5},{4,2,92,3,1},{1,2,1,5,89},{1,133,2,4,93},{5,142,95,5,92},{2,148,97,92,4},{97,142,98,4,1},{3,148,96,4,97},{96,148,1,97,2},{5,142,96,3,95},{1,133,5,5,89},{4,2,5,3,85},{5,1,90,3,1},{88,2,1,5,2},{94,133,96,4,4},{97,142,2,97,2},{2,148,97,94,4},{95,168,3,97,4},{3,196,98,96,5},{5,168,3,92,93},{3,142,5,4,87},{4,2,4,5,82},{1,1,3,97,95},{5,2,1,1,92},{91,2,1,133,5},{96,1,3,133,97},{1,2,94,3,1},{3,2,90,1,3},{2,2,3,92,2},{5,2,89,3,2},{4,2,2,1,89},{4,2,95,91,2},{1,2,90,5,2},{95,142,5,1,96},{2,168,96,96,5},{99,197,1,3,98},{2,197,99,1,95},{5,168,1,2,88},{4,142,95,96,5},{2,133,98,93,3},{4,2,2,89,2},{2,1,3,1,92},{4,2,96,5,92},{86,2,5,5,3},{1,1,92,3,2},{4,2,2,95,94},{94,1,97,2,3},{5,2,85,4,4},{89,2,2,2,4},{92,2,4,91,4},{94,131,4,1,98},{4,2,3,88,1},{1,131,1,98,98},{5,131,4,91,95},{2,2,91,4,2},{87,2,5,2,3},{1,2,5,86,2},{2,2,2,2,88},{4,1,2,86,5},{96,1,3,91,4},{95,2,1,2,1},{93,133,95,4,4},{2,142,4,92,96},{3,148,4,95,95},{2,142,1,1,95},{95,133,1,96,4},{3,1,5,96,96},{5,2,3,91,93},{5,2,96,92,1},{86,1,4,5,2},{96,2,3,2,94},{5,2,3,86,2},{2,1,1,2,92},{89,2,3,1,4},{1,2,95,1,2},{5,1,1,90,2},{4,2,5,82,5},{3,2,96,4,96},{5,133,5,95,94},{99,142,2,1,98},{5,148,91,1,3},{1,148,99,2,97},{96,142,1,1,1},{2,133,96,92,5},{2,1,4,92,94},{5,2,1,4,89},{3,2,2,86,4},{1,2,88,4,5},{4,2,90,2,4},{1,2,5,89,2},{4,2,1,87,5},{87,2,2,3,5},{4,2,98,1,96},{1,2,90,2,5},{92,2,3,4,1},{1,1,5,88,1},{1,1,2,3,93},{95,2,3,91,3},{4,1,2,92,96},{2,2,3,91,2},{2,1,2,89,4},{1,2,3,91,1},{3,2,95,93,4},{1,1,91,3,3},{94,1,1,3,97},{92,1,95,5,4},{3,1,2,96,94},{85,2,5,5,2},{5,131,95,4,95},{5,2,3,96,95},{5,2,88,4,2},{4,2,1,87,3},{5,2,4,2,86},{5,2,1,88,4},{4,2,3,92,1},{94,1,1,95,3},{1,2,2,5,89},{4,1,89,1,4},{2,2,91,3,4},{5,2,4,85,2},{3,1,2,92,1},{4,2,88,1,5},{5,2,95,3,95},{1,1,1,5,89},{96,2,96,3,4},{5,131,95,4,94},{4,131,1,95,97},{91,2,4,1,1},{88,2,2,3,4},{92,1,5,95,4},{4,1,92,1,2},{2,2,4,87,2},{95,1,4,1,97}}
//var res = []Resource{{1,2,3,4,5}}
func findMaxAndMix() {
	max := [5]int{}
	min := [5]int{100,100,100,100,100}
	for i := 0; i < len(res); i++{
		if res[i].gold > max[0] {
			max[0] = res[i].gold
		}
		if res[i].gold < min[0] {
			min[0] = res[i].gold
		}
		if res[i].wood > max[1] {
			max[1] = res[i].wood
		}
		if res[i].wood < min[1] {
			min[1] = res[i].wood
		}
		if res[i].water > max[2] {
			max[2] = res[i].water
		}
		if res[i].water < min[2] {
			min[2] = res[i].water
		}
		if res[i].fire > max[3] {
			max[3] = res[i].fire
		}
		if res[i].fire < min[3] {
			min[3] = res[i].fire
		}
		if res[i].earth > max[4] {
			max[4] = res[i].earth
		}
		if res[i].earth < min[4] {
			min[4] = res[i].earth
		}
	}
	fmt.Println(max,min)
}

func findDiffeent()  {
	diffGold := []int{}
	diffWood := []int{}
	diffWater := []int{}
	diffFire := []int{}
	diffEarth := []int{}


	diffCount := [5]int{}


	// go语言进行去重
	for i := 0; i < len(res); i++{
		flagGold := true
		for j := 0; j < len(diffGold); j++{
			if diffGold[j] == res[i].gold{
				flagGold = false // 存在重复
				break
			}
		}
		if flagGold {
			diffGold = append(diffGold,res[i].gold)
			diffCount[0]++
		}

		flagWood := true
		for j := 0; j < len(diffWood); j++{
			if diffWood[j] == res[i].wood{
				flagWood = false // 存在重复
				break
			}
		}
		if flagWood {
			diffWood = append(diffWood,res[i].wood)
			diffCount[1]++
		}

		flagWater := true
		for j := 0; j < len(diffWater); j++{
			if diffWater[j] == res[i].water{
				flagWater = false // 存在重复
				break
			}
		}
		if flagWater {
			diffWater = append(diffWater,res[i].water)
			diffCount[2]++
		}


		flagFire := true
		for j := 0; j < len(diffFire); j++{
			if diffFire[j] == res[i].fire{
				flagFire = false // 存在重复
				break
			}
		}
		if flagFire {
			diffFire = append(diffFire,res[i].fire)
			diffCount[3]++
		}



		flagEarth := true
		for j := 0; j < len(diffEarth); j++{
			if diffEarth[j] == res[i].earth{
				flagEarth = false // 存在重复
				break
			}
		}
		if flagEarth {
			diffEarth = append(diffEarth,res[i].earth)
			diffCount[4]++
		}

	}
	sort.Ints(diffGold)
	fmt.Println(diffGold)
	//fmt.Println(diffWood)
	//fmt.Println(diffWater)
	fmt.Println(diffCount)
}


func main()  {
	findMaxAndMix()
	findDiffeent()
}