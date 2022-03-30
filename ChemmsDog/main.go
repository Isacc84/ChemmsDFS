package main

import (
	"bufio"
	"container/list"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/shirou/gopsutil/v3/disk"
)

var (
	talkArr         chan []uint8
	messageTypeKage chan int
	DisList         list.List
	tmpDiskse       string
	ms              []uint8
	er              error
	er0             error
	messageTyp0     int
	p0              []byte
	Pussy           *uint64
)

var tmpDiskset = &tmpDiskse
var noUseArr sync.Map
var msg = &ms
var err = &er
var te string
var tem = &te
var messageType = &messageTyp0
var p1 = &p0
var err1 = &er0
var err0 = &er
var pussyCap sync.Map
var che int
var chae = &che
var Msttru bool
var Slatru bool
var MSt = &Msttru
var Sla = &Slatru
var allStau sync.Map
var returnLis = make([]string, 0, 10000)
var returnLisKaGe = make([]uint64, 0, 10000)
var onlis = make([]string, 0, 10000)
var nolis = make([]string, 0, 10000)
var slaveBaby = make([]string, 0, 10000)
var cha chan sync.Map

func DiskFreeMemory() (*[]string, *[]uint64, *sync.Map, error) {
	mapStat, _ := disk.IOCounters()
	for name, _ := range mapStat {
		DisList.PushBack(name)
	}
	// 获取磁盘信息后输入到两个列表中
	for item := DisList.Front(); item != nil; item = item.Next() {
		info0, _ := disk.Usage(item.Value.(string))
		// data0, _ := json.MarshalIndent(info0, "", "  ")
		// fmt.Println(info0)
		//如果每一项的值不是空那就循环，不然就读取失败
		if reflect.TypeOf(info0) != nil {
			val, ok := pussyCap.Load(info0.Path)
			if val == nil || ok == false {
				pussyCap.Store(info0.Path, *chae)
				returnLis = append(returnLis, info0.Path)
				returnLisKaGe = append(returnLisKaGe, info0.Free)
				*chae += 1
			} else {
				for k, _ := range returnLisKaGe {
					if k == *chae {
						returnLisKaGe[k] = info0.Free
					}
				}
			}
		} else {
			err := errors.New("读取磁盘大失败")
			return nil, nil, nil, err
		}
	}
	var lisRan uint64 = uint64(len(returnLis))
	var lisRanKaGe uint64 = uint64(len(returnLisKaGe))
	if lisRan != lisRanKaGe {
		err := errors.New("获取硬盘列表失败")
		return nil, nil, nil, err
	}
	for keys := lisRanKaGe - 1; uint64(keys) > 0; keys-- {
		for key := lisRanKaGe - 1; uint64(key) > 0; key-- {
			if returnLisKaGe[key] >= returnLisKaGe[key-1] {
				temCup := returnLisKaGe[key]
				returnLisKaGe[key] = returnLisKaGe[key-1]
				returnLisKaGe[key-1] = temCup
				temCup0 := returnLis[key]
				returnLis[key] = returnLis[key-1]
				returnLis[key-1] = temCup0
			}
		}
	}
	//对列表内容进行了排序
	var mapDikMes sync.Map
	for keys := 0; uint64(keys) < lisRan; keys++ {
		mapDikMes.Store(returnLis[keys], returnLisKaGe[keys])
	}
	return &returnLis, &returnLisKaGe, &mapDikMes, nil
}
func procoser(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	// fmt.Println("读取等")
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("读取失败")
		for k, v := range onlis {
			fmt.Println(conn.RemoteAddr(), string(buf[:n]), k, v)
			if v == strings.Split(conn.RemoteAddr().String(), ":")[0] {
				onlis[k] = "none"
			}
			// if strings.Split(v,":")[0]
		}
		return
	}
	var infos int64
	for k, v := range onlis {
		fmt.Println(conn.RemoteAddr(), string(buf[:n]), k, v)
		if strings.Split(v, ":")[0] == strings.Split(conn.RemoteAddr().String(), ":")[0] {
			infos += 1
			onlis[k] = strings.Split(conn.RemoteAddr().String(), ":")[0] + ":" + string(buf[:n])
		}
		// if strings.Split(v,":")[0]
	}
	if infos == 0 {
		onlis = append(onlis, strings.Split(conn.RemoteAddr().String(), ":")[0]+":"+string(buf[:n]))
	}
	allStau.Store("online", onlis)
	fmt.Println("关庙链接", conn.RemoteAddr(), string(buf[:n]), len(onlis), onlis)
}
func refun() {
	fileInfo, _ := os.Lstat("./option/CJokerSet/SlaveServer.csv")
	if fileInfo == nil {
		fmt.Println("缺少子树文件./option/CJokerSet/SlaveServer.csv")
	} else {
		file, err := os.OpenFile("./option/CJokerSet/SlaveServer.csv", os.O_RDONLY, 0777)
		defer file.Close()
		if err != nil {
			// fmt.Println("打开文件问题")
		} else {
			rd := bufio.NewReader(file)
			for {
				line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
				if line == "" || line == "\n" {
					*Sla = true
					*MSt = false
					fmt.Println("这是最低节点不再监听1810端口①", onlis)
					break
				}
				if err != nil || io.EOF == err {
					break
				}
			}
			if *Sla == true && *MSt == false {
				return
			} else {
				fmt.Println("开始监听")
				liserm, err := net.Listen("tcp", "0.0.0.0:1810")
				if err != nil {
					fmt.Println(err)
					return
				}
				defer liserm.Close()
				for {
					conn, err := liserm.Accept()
					if err != nil {
						fmt.Println("监听失败的原因", err)
						for k, v := range onlis {
							if strings.Split(v, ":")[0] == conn.RemoteAddr().String() {
								onlis[k] = "none"
							}
						}
					} else {
						go procoser(conn)
						fmt.Printf("监听成功%T%v\n", conn.RemoteAddr().String(), conn.RemoteAddr().String())
					}
				}
			}
		}
	}
}

func loop() {
	time.Sleep(time.Second)
	time.Sleep(time.Second)
	time.Sleep(time.Second)
	time.Sleep(time.Second)
	time.Sleep(time.Second)
	triememo()
	allStau.Range(func(k, v interface{}) bool {
		fmt.Println(k, v)
		return true
	})
	cha <- allStau
}

var (
	memory     int64
	triememory int64
)

func triememo() {
	memory = 0
	triememory = 0
	a, b, _, _ := DiskFreeMemory()
	aa := *a
	file, err := os.OpenFile("./option/CJokerSet/CDFSsetfile.csv", os.O_RDONLY, 0777)
	defer file.Close()
	if err != nil {
		fmt.Println("打开fei文件问题")
		return
	} else {
		rd := bufio.NewReader(file)
		for {
			tmpArr, err := rd.ReadString('\n')
			asa := strings.Split(tmpArr, ",")
			for _, v := range asa {
				for ks, vs := range *b {
					if aa[ks] == v {
						memory += int64(vs) / (1024 * 1024 * 1024)
						allStau.Store(aa[ks], int64(vs)/(1024*1024*1024))
					}
				}
			}
			if err != nil || io.EOF == err {
				break
			}
		}
		allStau.Store("user", int64(memory))
	}
	for _, vlu := range onlis {
		dis, _ := strconv.ParseInt(strings.Split(vlu, ":")[1], 10, 64)
		triememory += dis
	}
	allStau.Store("tri", int64(triememory))
	fmt.Println("sahdjh", memory, triememory)
	fileInfo, _ := os.Lstat("./option/CJokerSet/MaseterServer.csv")
	if fileInfo == nil {
		fmt.Println("缺少主机文件./option/CJokerSet/MaseterServer.csv")
	} else {
		file, err := os.OpenFile("./option/CJokerSet/MaseterServer.csv", os.O_RDONLY, 0777)
		defer file.Close()
		if err != nil {
			fmt.Println("打开文件问题")
			return
		} else {
			rd := bufio.NewReader(file)
			for {
				line, err := rd.ReadString('\n')
				if line == "" {
					fmt.Println("这主机")
					break
				}
				conn, err := net.Dial("tcp", line+":1810")
				defer conn.Close()
				if err != nil {
					fmt.Println("连接错误，该主机可能宕机，介入备份主", err)
					continue
				}
				_, err0 := conn.Write([]byte(strconv.FormatInt(memory+triememory, 10)))
				if err0 != nil {
					fmt.Println("一般般", err0)
				} else {
					fmt.Println("一翻身")
				}
				if err != nil || io.EOF == err {
					break
				}
			}
		}
	}
	go loop()
}

//风格线，下面都是CDFS
//风格线，下面都是CDFS
//风格线，下面都是CDFS
//风格线，下面都是CDFS
//风格线，下面都是CDFS
//风格线，下面都是CDFS
var a0 sync.Map

var addRFi []string
var chaZiDian map[string]interface{}

func MapNewLop(cha *chan sync.Map) {
	time.Sleep(time.Second)
	time.Sleep(time.Second)
	time.Sleep(time.Second)
	time.Sleep(time.Second)
	time.Sleep(time.Second)
	go MapNewStu(cha)
}

var slaAr []string
var slaLr []int64
var slaAar chan []string
var slaLar chan []int64

func MapNewStu(cha *chan sync.Map) {
	slaAr = make([]string, 100000)
	slaLr = make([]int64, 100000)
	fmt.Println(len(*cha))
	a0 = allStau
	a0.Range(func(k, v interface{}) bool {
		fmt.Println(k, v)
		if k.(string) == "online" {
			fmt.Println("hduiawhusz", v)
			for k0, i := range v.([]string) {
				slaAr[k0] = strings.Split(i, ":")[0]
				tek, _ := strconv.ParseInt(strings.Split(i, ":")[1], 10, 64)
				slaLr[k0] = tek
				//通过分割数据，获得每个服务器应该传几个文件，并在切分文件的时候就将文件写入对应零食文件夹，然后遍历文件进行上传对应服务
			}
		}
		return true
	})
	var lisRan int64 = int64(len(slaAr))
	var lisRanKaGe int64 = int64(len(slaLr))
	if lisRan != lisRanKaGe {
		fmt.Println("失败零时文件")
		return
	}
	// for keys := lisRanKaGe - 1; int64(keys) > 0; keys-- {
	// 	for key := lisRanKaGe - 1; int64(key) > 0; key-- {
	// 		if slaLr[key] >= slaLr[key-1] {
	// 			temCup := slaLr[key]
	// 			slaLr[key] = slaLr[key-1]
	// 			slaLr[key-1] = temCup
	// 			temCup0 := slaAr[key]
	// 			slaAr[key] = slaAr[key-1]
	// 			slaAr[key-1] = temCup0
	// 		}
	// 	}
	// }
	slaAar <- slaAr
	slaLar <- slaLr
	// fmt.Println(slaLr, slaAr)
	go MapNewLop(cha)
}
func writeLop() {
	go writeLi()
}

func writeLi() {
	// var tmPapS string
	if len(brautyPussy) > 0 {
		pp := <-girlPussy
		np := <-brautyPussy
		stc := <-graggyPussy
		if np == 810 {
			fileInfo, err := os.Lstat(stc)
			fmt.Println(err)
			if fileInfo == nil {
				file, err := os.OpenFile(stc, os.O_WRONLY|os.O_CREATE, 0777)
				defer file.Close()
				if err != nil {
					fmt.Println(err)
				} else {
					writer := bufio.NewWriter(file)
					writer.WriteString(string(pp))
					writer.Flush()
				}
				file.Close()
			} else {
				fmt.Println("文件存在谢绝放鸽子1")
			}
		}
		// if np == 1919 {
		// 	file, err := os.OpenFile(stc, os.O_WRONLY|os.O_CREATE, 0777)
		// 	defer file.Close()
		// 	if err != nil {
		// 		fmt.Println(err)
		// 	} else {
		// 		writer := bufio.NewWriter(file)
		// 		writer.WriteString(string(pp))
		// 		writer.Flush()
		// 	}
		// 	file.Close()
		// 	// getZawarodoToMadeinheaven(stc)
		// 	continue
		// }
		if string(pp[:np]) == "" {
			fmt.Println("空字符串了已")
			st <- 1
			fmt.Println()
			fmt.Println()
			fmt.Println()
			fmt.Println()
			fmt.Println(len(st))
			go writeLop()
			return
		}
		// fmt.Printf("%-20s 这就是%-2v %v\n", pp, np, stc)
		if np != 810 {
			fileInfo, err := os.Lstat(stc)
			fmt.Println(err)
			if fileInfo == nil {
				file, err := os.OpenFile(stc, os.O_WRONLY|os.O_CREATE, 0777)
				defer file.Close()
				if err != nil {
					fmt.Println(err)
				} else {
					writer := bufio.NewWriter(file)
					writer.WriteString(string(pp))
					writer.Flush()
				}
				file.Close()
			} else {
				fmt.Println("文件存在谢绝放鸽子2")
			}
		}
	}
	go writeLop()
	return
}

var girlPussy chan []byte
var brautyPussy chan int
var graggyPussy chan string
var st chan int

var CHesm1 chan string
var CHesm2 chan string
var CHesm3 chan string

func deLop() {
	go delFil()
}

func delFil() {
	//建立对每个服务器的连接，发送文件，之后对面收,改管道，一次只发一个部分
	if len(CHesm3) > 0 {
		fmt.Println("建立对每个服务器的连接，发送文件，之后对面收,改管道，一次只发一个部分")
		slaAr := <-CHesm1
		cpa := <-CHesm2
		fhu := <-CHesm3
		conn, err := net.Dial("tcp", slaAr+":1145")
		fmt.Println(slaAr)
		if err != nil {
			fmt.Println("连接错误", err)
			return
		}
		defer conn.Close()
		file, err := os.OpenFile(cpa, os.O_RDONLY, 0777)
		if err != nil {
			fmt.Println("读取失败", err)
		}
		fife, _ := os.Stat(cpa)
		var DMT = []byte("|||" + fhu)
		br := bufio.NewReaderSize(file, int(fife.Size()))
		defer file.Close()
		p := make([]byte, int(fife.Size()))
		n, _ := br.Read(p)
		fmt.Println(n)
		for _, v := range DMT {
			p = append(p, v)
		}
		_, err = conn.Write(p)
		if err != nil {
			fmt.Println("发送大失败了", err)
		} else {
			file.Close()
			for k, v := range addRFi {
				if strings.Split(v, ":")[0] == "onSlave" {
					if strings.Split(v, ":")[1] == fhu {
						addRFi[k] = "onServ:" + strings.Split(conn.RemoteAddr().String(), ":")[0] + "|" + fhu
						fmt.Println("onServ:" + strings.Split(conn.RemoteAddr().String(), ":")[0] + "|" + fhu)
						chaZiDian[tmpSty] = addRFi
					}
				}
			}
			err := os.Remove(cpa)
			if err != nil {
				fmt.Println("删除零时文件失败", err)
			}
			metaSVE()
		}
		conn.Close()
	}
	go deLop()
}

func getZawarodoToMadeinheaven() {
	if len(slaLar) >= 1 {
		slaLr := <-slaLar
		slaAr := <-slaAar
		fmt.Println("取值成功")
		var cpa = make([]string, 0, 100000)
		var fhu = make([]string, 0, 100000)
		a0.Range(func(k, v interface{}) bool {
			if k != "all" && k != "online" && k != "user" && k != "tri" {
				temMes, err := ioutil.ReadDir(k.(string) + "/CDFS/temPreegPussy/")
				if err != nil {
					fmt.Print("分配错误")
					fmt.Println(err)
				}
				for _, file := range temMes {
					if file.Name() == "" {
						break
					}
					fmt.Println(file.Name())
					fmt.Printf("%T", file.Name())
					fhu = append(fhu, file.Name())
					cpa = append(cpa, k.(string)+"/CDFS/temPreegPussy/"+file.Name())
				}
				//通过将文件嘉进行遍历，获得每个文件尺寸，根据仆从鸡内部大小安排文件
			}
			return true
		})
		var lop = len(cpa)
		var dps int64
		fmt.Println("有", lop)
		for lop > 0 {
			for k, v := range slaLr {
				if slaLr[0] > 0 {
					if v <= 0 && k != 0 && lop > 0 {
						continue
					} else {
						filInf, _ := os.Lstat(cpa[dps])
						if filInf != nil {
							fmt.Println(1)
							CHesm1 <- slaAr[k]
							CHesm2 <- cpa[dps]
							CHesm3 <- fhu[dps]
							slaLr[k] -= 1
							lop -= 1
							dps += 1
						} else {
							dps += 1
						}
						if lop <= 0 {
							fmt.Println("超过内容")
							return
						}
					}
				} else {
					break
				}
			}
		}
	}
	go betozadasuto()
}

func betozadasuto() {
	go getZawarodoToMadeinheaven()
}

func filEat(AdL string, AdL0 string) {
	fmt.Println("开始")
	addRFi = make([]string, 0, 100000)
	fmt.Println(len(girlPussy))
	fmt.Println(len(brautyPussy))
	fmt.Println(len(graggyPussy))
	fmt.Println()
	fi, e := os.Stat(AdL)
	var fis int64
	if fi != nil || e == nil {
		fis = int64(fi.Size())
	} else {
		fmt.Println("无法获取文件大小")
		return
	}
	file, err := os.OpenFile(AdL, os.O_RDONLY, 0777)
	if err != nil {
		fmt.Println("读取失败", err)
	}
	defer file.Close()
	fife, _ := os.Stat(AdL)
	br := bufio.NewReaderSize(file, int(fife.Size()))
	var nas uint64
	for len(st) < 1 {
		a0.Range(func(k, v interface{}) bool {
			if k != "all" && k != "online" && k != "user" && k != "tri" {
				fis -= v.(int64) * 1024 * 50
				if fis < 0 {
					v = v.(int64) - (fis + v.(int64))
					for itr := 0; int64(itr) < v.(int64); itr++ {
						p := make([]byte, 1024*50)
						n, _ := br.Read(p)
						var tes = k.(string) + "/CDFS/hostPreegPussy/" + AdL0 + "_" + strconv.FormatUint(nas, 10)
						graggyPussy <- tes
						girlPussy <- p
						brautyPussy <- n
						addRFi = append(addRFi, tes)
						chaZiDian[AdL0] = addRFi
						// fmt.Println("a", graggyPussy, brautyPussy, girlPussy)
						nas += 1
					}
					return true
				} else {
					for itr := 0; int64(itr) < v.(int64); itr++ {
						p := make([]byte, 1024*50)
						n, _ := br.Read(p)
						var tes = k.(string) + "/CDFS/hostPreegPussy/" + AdL0 + "_" + strconv.FormatUint(nas, 10)
						graggyPussy <- tes
						girlPussy <- p
						brautyPussy <- n
						addRFi = append(addRFi, tes)
						chaZiDian[AdL0] = addRFi
						// fmt.Println("b", graggyPussy, brautyPussy, girlPussy)
						nas += 1
					}
				}
			}
			if k == "online" {
				a0.Range(func(k, v interface{}) bool {
					if k != "all" && k != "online" && k != "user" && k != "tri" {
						fis -= v.(int64) * 1024 * 50
						if fis < 0 {
							v = v.(int64) - (fis + v.(int64))
							for itr := 0; int64(itr) < v.(int64); itr++ {
								p := make([]byte, 1024*50)
								n, _ := br.Read(p)
								var tes = k.(string) + "/CDFS/temPreegPussy/" + AdL0 + "_" + strconv.FormatUint(nas, 10)
								graggyPussy <- tes
								girlPussy <- p
								brautyPussy <- n
								fmt.Println("王道征途人", v, itr)
								addRFi = append(addRFi, "onSlave:"+AdL0+"_"+strconv.FormatUint(nas, 10))
								chaZiDian[AdL0] = addRFi
								nas += 1
								// fmt.Println("a", graggyPussy, brautyPussy, girlPussy)
							}
							return true
						} else {
							for itr := 0; int64(itr) < v.(int64); itr++ {
								p := make([]byte, 1024*50)
								n, _ := br.Read(p)
								var tes = k.(string) + "/CDFS/temPreegPussy/" + AdL0 + "_" + strconv.FormatUint(nas, 10)
								graggyPussy <- tes
								girlPussy <- p
								brautyPussy <- n
								fmt.Println("野兽昏睡茶")
								addRFi = append(addRFi, "onSlave:"+AdL0+"_"+strconv.FormatUint(nas, 10))
								chaZiDian[AdL0] = addRFi
								nas += 1
								// fmt.Println("b", graggyPussy, brautyPussy, girlPussy)
							}
							return true
						}
					}
					return true
				})
			}
			fmt.Println()
			fmt.Println()
			fmt.Println()
			fmt.Println("结束")
			metaSVE()
			return true
		})
	}
	if len(st) > 0 {
		ass := <-st
		fmt.Println("退写文件循环", ass)
	}
}

func sliEat(AdL string, AdL0 string) {
	fmt.Println("开始")
	fmt.Println(len(girlPussy))
	fmt.Println(len(brautyPussy))
	fmt.Println(len(graggyPussy))
	fmt.Println()
	addRFi = make([]string, 0, 100000)
	file, err := os.OpenFile(AdL, os.O_RDONLY, 0777)
	if err != nil {
		fmt.Println("读取失败", err)
	}
	defer file.Close()
	br := bufio.NewReaderSize(file, 1024*1024)
	var nas uint64
	for len(st) < 1 {
		a0.Range(func(k, v interface{}) bool {
			if k != "all" && k != "online" && k != "user" && k != "tri" {
				for itr := 0; int64(itr) < v.(int64); itr++ {
					p, _ := br.ReadBytes('\n')
					suck := len(p)
					for suck <= 1024*50 {
						pp, _ := br.ReadBytes('\n')
						suck += len(pp)
						for _, v := range pp {
							p = append(p, v)
						}
					}
					n := 810
					var tes = k.(string) + "/CDFS/hostPreegPussy/" + AdL0 + "_" + strconv.FormatUint(nas, 10)
					graggyPussy <- tes
					girlPussy <- p
					brautyPussy <- n
					fmt.Println("野兽昏睡茶", v, itr)
					addRFi = append(addRFi, tes)
					chaZiDian[AdL0] = addRFi
					// fmt.Println("a", graggyPussy, brautyPussy, girlPussy)
					fmt.Println(addRFi)
					nas += 1
				}
				metaSVE()
				return true
			}
			if k == "online" {
				a0.Range(func(k, v interface{}) bool {
					if k != "all" && k != "online" && k != "user" && k != "tri" {
						for itr := 0; int64(itr) < v.(int64); itr++ {
							p, _ := br.ReadBytes('\n')
							suck := len(p)
							for suck <= 1024*50 {
								pp, _ := br.ReadBytes('\n')
								suck += len(pp)
								for _, v := range pp {
									p = append(p, v)
								}
							}
							n := 810
							var tes = k.(string) + "/CDFS/temPreegPussy/" + AdL0 + "_" + strconv.FormatUint(nas, 10)
							graggyPussy <- tes
							girlPussy <- p
							brautyPussy <- n
							fmt.Println("王道征途人", v, itr)
							addRFi = append(addRFi, "onSlave:"+AdL0+"_"+strconv.FormatUint(nas, 10))
							chaZiDian[AdL0] = addRFi
							fmt.Println(addRFi)
							nas += 1
							// fmt.Println("a", graggyPussy, brautyPussy, girlPussy)
						}
					}
					return true
				})
			}
			fmt.Println()
			fmt.Println()
			fmt.Println()
			fmt.Println("结束")
			metaSVE()
			return true
		})
	}
	if len(st) > 0 {
		ass := <-st
		fmt.Println("退写文件循环", ass)
	}
}

var PreggyPussy = make([]string, 100, 100000)
var boobSize uint64
var tit = &boobSize

func sopt() {
	fileInfo, _ := os.Lstat("./option/CJokerSet/CDFSsetfile.csv")
	if fileInfo == nil {
		fmt.Println("./option/CJokerSet/CDFSsetfile.txt不存在")
	} else {
		file, err := os.OpenFile("./option/CJokerSet/CDFSsetfile.csv", os.O_RDONLY, 0777)
		if err != nil {
			fmt.Println("无法打不进文件", err)
		}
		defer file.Close()
		rli := bufio.NewReader(file)
		for {
			line, err := rli.ReadString('\n') //以'\n'为结束符读入一行
			if err != nil || io.EOF == err {
				break
			}
			if strings.Split(line, ",")[0] == "UseDiskLis" {
				for k, v := range strings.Split(line, ",") {
					if k == 0 || v == "\n" {
						continue
					}
					err := os.Mkdir(v+"/CDFS/", 0777)
					PreggyPussy = append(PreggyPussy, v+"/CDFS/")
					if err != nil {
						fmt.Println("创建文件失败", err, v+"/CDFS/")
					}
					os.Mkdir(v+"/CDFS/hostPreegPussy", 0777)
					os.Mkdir(v+"/CDFS/slavePreegPussy", 0777)
					os.Mkdir(v+"/CDFS/temPreegPussy", 0777)
				}
			}
			if strings.Split(line, ",")[0] == "boobsize" {
				*tit, err = strconv.ParseUint(strings.Split(line, ",")[1], 10, 64)
				if err != nil {
					fmt.Println("错误", err)
				}
			}
		}
	}
}

var thoren int64
var metaSlave []string

func procoser114(conn net.Conn) {
	thoren += 1
	fmt.Println(thoren)
	defer conn.Close()
	buf := make([]byte, 1024*1024)
	var returns = 0
	for returns < 1 {
		if returns >= 1 {
			break
		}
		a0.Range(func(k, _ interface{}) bool {
			if k != "all" && k != "online" && k != "user" && k != "tri" {
				// fmt.Println("读取等")
				n, err := conn.Read(buf)
				if err != nil {
					fmt.Println("读取错误子机可能关闭了连接", err)
					returns += 1
					return false
				} else {
					sts := string(buf[:n])
					stA := strings.Split(sts, "|||")
					// fmt.Println(stA)
					if len(stA) <= 1 {
						fmt.Println("文件缺少")
						return true
					}
					file, err := os.OpenFile(k.(string)+"/CDFS/slavePreegPussy/"+stA[1], os.O_CREATE|os.O_WRONLY, 0777)
					defer file.Close()
					// fmt.Println(k.(string) + "/CDFS/slavePreegPussy/" + stA[1])
					if err != nil {
						fmt.Print("创建件大失败")
						return true
					} else {
						writer := bufio.NewWriter(file)
						writer.WriteString(stA[0])
						writer.Flush()
						metaSlave = append(metaSlave, k.(string)+"/CDFS/slavePreegPussy/"+stA[1])
						chaZiDian[strings.Split(stA[1], "_")[0]] = metaSlave
						metaSVE()
					}
				}
			}
			return true
		})
	}
	thoren -= 1
	fmt.Println("释放", thoren)
}

func TcpSAO0() {
	metaSlave = make([]string, 0, 100000)
	fmt.Println("服务器开始监听")
	liserm, err := net.Listen("tcp", "0.0.0.0:1145")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer liserm.Close()
	for {
		conn, err := liserm.Accept()
		if err != nil {
			fmt.Println("读问题")
		} else {
			fmt.Printf("监听成功%T%v\n", conn.RemoteAddr().String(), conn.RemoteAddr().String())
			go procoser114(conn)
		}
	}
	fmt.Println("不减停")
}
func metaSVE() {
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("运行")
	data, err := json.Marshal(chaZiDian)
	if err != nil {
		fmt.Println("元数据无职转生出错")
	}
	file, err := os.OpenFile("./meta/"+tmpSty+".meta", os.O_WRONLY|os.O_CREATE, 0777)
	writer := bufio.NewWriter(file)
	writer.WriteString(string(data))
	writer.Flush()
	file.Close()
	fmt.Println("运行了")
}

var tmpSty = "Bcard"

func procoser0(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024*1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("打错", err, n)
	}
	sts := string(buf[:n])
	fmt.Println(sts)
	if strings.Split(sts, "|")[0] == "u" {
		if strings.Split(sts, "|")[1] == "for" {
			fmt.Println("had岁啊和我的", strings.Split(sts, "|"))
			sliEat("./upload/Bcard.txt", "Bcard")
		}
		if strings.Split(sts, "|")[1] == "nor" {
			filEat("./upload/Bcard.txt", "Bcard")
		}
	}
	if strings.Split(sts, "|")[0] == "o" {
		fmt.Println(strings.Split(sts, "|")[4])
		isdohfgjhs()
	}

}

func TcpSAO1() {
	fmt.Println("服务器开始监听")
	liserm, err := net.Listen("tcp", "0.0.0.0:5140")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer liserm.Close()
	for {
		conn, err := liserm.Accept()
		if err != nil {
			fmt.Println("读问题")
		} else {
			fmt.Printf("监听成功%T%v\n", conn.RemoteAddr().String(), conn.RemoteAddr().String())
			go procoser0(conn)
		}
	}
	fmt.Println("不减停")
}
func Cofile(tarFil string, getFil string) (wrien int64, err error) {
	srcFil, err := os.Open(getFil)
	if err != nil {
		fmt.Printf("openerr&v\n", err)
	}
	defer srcFil.Close()
	reader := bufio.NewReader(srcFil)
	dstFil, err := os.OpenFile(tarFil, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		fmt.Printf("openerr&v\n", err)
	}
	writer := bufio.NewWriter(dstFil)
	defer dstFil.Close()
	return io.Copy(writer, reader)
}

var temSts string

var tmpMod map[string]interface{}

var sts string

func isdohfgjhs() {
	file, err := os.OpenFile("./meta/Bcard.meta", os.O_RDONLY, 0777)
	if err != nil {
		fmt.Println("元数据鼠标", err)
	}
	br := bufio.NewReader(file)
	fil, err := br.ReadBytes('}')
	if err != nil {
		fmt.Println("读取大师班")
	}
	errs := json.Unmarshal(fil, &tmpMod)
	if errs != nil {
		fmt.Println("转反序列失败")
	}
	nas, ok := tmpMod[strings.Split(sts, "|")[4]]
	fmt.Println(tmpMod)
	fmt.Println(strings.Split(sts, "|")[4])
	if ok {
		for _, v := range nas.([]interface{}) {
			if strings.Split(v.(string), ":")[0] == "onServ" {
				if temSts == strings.Split(strings.Split(v.(string), ":")[1], "|")[0] {
					continue
				}
				conn, err := net.Dial("tcp", strings.Split(strings.Split(v.(string), ":")[1], "|")[0]+":1949")
				if err != nil {
					fmt.Println("连接错误", err)
				}
				n, errLei := conn.Write([]byte(strings.Split(sts, "|")[4]))
				if errLei != nil {
					fmt.Println("发送大失败了", err, n)
				} else {
					temSts = strings.Split(strings.Split(v.(string), ":")[1], "|")[0]
				}
				defer conn.Close()
			} else {
				Cofile("./temDir/"+strings.Split(v.(string), "/")[len(strings.Split(v.(string), "/"))-1], v.(string))
			}
		}
	} else {
		fmt.Println("我这个不到啊")
	}
}

func procoser01(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024*1024)
	n, err := conn.Read(buf)
	fmt.Println(string(buf[:n]), 1)
	if err != nil {
		fmt.Println("打错", err, n)
	}
	file, err := os.OpenFile("./meta/Bcard.meta", os.O_RDONLY, 0777)
	if err != nil {
		fmt.Println("元数据鼠标", err)
	}
	br := bufio.NewReader(file)
	fil, err := br.ReadBytes('}')
	if err != nil {
		fmt.Println("读取大师班")
	}
	var tmpMod map[string]interface{}
	errs := json.Unmarshal(fil, &tmpMod)
	if errs != nil {
		fmt.Println("转反序列失败")
	}
	nas, ok := tmpMod[string(buf[:n])]
	fmt.Println(nas, string(buf[:n]), 2, tmpMod)
	if ok {
		for _, v := range nas.([]interface{}) {
			if strings.Split(v.(string), ":")[0] == "onServ" {
				continue
			} else {
				fmt.Println(strings.Split(conn.RemoteAddr().String(), ":")[0], 3)
				connz, err := net.Dial("tcp", strings.Split(conn.RemoteAddr().String(), ":")[0]+":1927")
				if err != nil {
					fmt.Println("连接错误", err)
					continue
				}
				file, err := os.OpenFile(v.(string), os.O_RDONLY, 0777)
				fife, _ := os.Stat(v.(string))
				var boobs = make([]byte, int(fife.Size()+100))
				if err != nil {
					fmt.Println("元数据鼠标", err)
				}
				br := bufio.NewReader(file)
				dickLen, err := br.Read(boobs)
				if err != nil {
					fmt.Println("读取大师班")
				}
				sys := string(boobs[:dickLen]) + "|||" + strings.Split(v.(string), "/")[len(strings.Split(v.(string), "/"))-1]
				n, errLei := connz.Write([]byte(sys))
				if errLei != nil {
					fmt.Println("发送大失败了", err, n)
				}
				defer connz.Close()
			}
		}
	}
}

func TcpSAO10() {
	fmt.Println("服务器开始监听")
	liserm, err := net.Listen("tcp", "0.0.0.0:1949")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer liserm.Close()
	for {
		conn, err := liserm.Accept()
		if err != nil {
			fmt.Println("读问题")
		} else {
			fmt.Printf("监听成功%T%v\n", conn.RemoteAddr().String(), conn.RemoteAddr().String())
			go procoser01(conn)
		}
	}
	fmt.Println("不减停")
}

func procoser010(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024*1024*10)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("打错", err, n)
	}
	nas := string(buf[:n])
	var tmpNsa int64
	var tmeSqs string
	if len(strings.Split(nas, "|||")) > 1 {
		file, err := os.OpenFile("./temDir/"+strings.Split(nas, "|||")[1], os.O_CREATE|os.O_WRONLY, 0777)
		defer file.Close()
		// fmt.Println(k.(string) + "/CDFS/slavePreegPussy/" + stA[1])
		if err != nil {
			fmt.Print("创建件大失败")
			return
		} else {
			writer := bufio.NewWriter(file)
			writer.WriteString(strings.Split(nas, "|||")[0])
			writer.Flush()
		}
		tmpNsa, _ = strconv.ParseInt(strings.Split(strings.Split(nas, "|||")[1], "_")[1], 10, 64)
		tmeSqs = strings.Split(strings.Split(nas, "|||")[1], "_")[0]
	} else {
		fmt.Println("文件残缺了自動補全")
		file, err := os.OpenFile("./temDir/"+tmeSqs+"_"+strconv.FormatInt(tmpNsa, 10), os.O_CREATE|os.O_WRONLY, 0777)
		defer file.Close()
		// fmt.Println(k.(string) + "/CDFS/slavePreegPussy/" + stA[1])
		if err != nil {
			fmt.Print("创建件大失败")
			return
		} else {
			writer := bufio.NewWriter(file)
			writer.WriteString(nas)
			writer.Flush()
		}
	}
	fmt.Println("释放")
}

var copWod chan string
var copWod0 chan string

func TcpSAO101() {
	fmt.Println("服务器开始监听")
	liserm, err := net.Listen("tcp", "0.0.0.0:1927")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer liserm.Close()
	for {
		conn, err := liserm.Accept()
		if err != nil {
			fmt.Println("读问题")
		} else {
			fmt.Printf("监听成功%T%v\n", conn.RemoteAddr().String(), conn.RemoteAddr().String())
			go procoser010(conn)
		}
	}
	fmt.Println("不减停")
}

func main() {
	cha = make(chan sync.Map, 100)
	go refun()
	go loop()
	copWod = make(chan string, 100000)
	copWod0 = make(chan string, 100000)
	file, err := os.OpenFile("./meta/Bcard.meta", os.O_RDONLY, 0777)
	if err != nil {
		fmt.Println("元数据鼠标", err)
	}
	br := bufio.NewReader(file)
	fil, err := br.ReadBytes('}')
	if err != nil {
		fmt.Println("读取大师班")
	}
	fmt.Println(string(fil))
	errs := json.Unmarshal(fil, &tmpMod)
	if errs != nil {
		fmt.Printf("转反序列失败%v\n", errs)
	}
	fmt.Println(tmpMod)
	go TcpSAO10()
	go TcpSAO101()
	chaZiDian = make(map[string]interface{})
	girlPussy = make(chan []byte, 1000)
	brautyPussy = make(chan int, 1000)
	graggyPussy = make(chan string, 1000)
	CHesm1 = make(chan string, 1000)
	CHesm2 = make(chan string, 1000)
	CHesm3 = make(chan string, 1000)
	slaLar = make(chan []int64, 1000)
	slaAar = make(chan []string, 1000)
	st = make(chan int, 1)
	sopt()
	go writeLi()
	go TcpSAO0()
	go delFil()
	go MapNewStu(&cha)
	go getZawarodoToMadeinheaven()
	go TcpSAO1()
	var a uint64
	var Pussy = &a
	for {
		fmt.Print("卡其脱离太", *Pussy, "分钟")
		time.Sleep(time.Minute)
		*Pussy = *Pussy + 1
	}
}
