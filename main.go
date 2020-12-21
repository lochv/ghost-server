package main

import (
	"github.com/kataras/iris/v12"
	"os"
	"fmt"
	"bufio"
	"strconv"
	"strings"
	"io/ioutil"
)


func main() {
	// RestoreAssets("","static")
    app := iris.Default()
    app.Get("/toolv3/home.php", func(ctx iris.Context) {
        ctx.Writef(`<body ondragstart="return false" ondragover="return false"><style type="text/css">
		    * {
		        -webkit-user-select: none;
		        cursor: default;   
		    }      
		    .btn{

		        border-radius: 2px;
		        box-shadow: 0 0 3px #333;
		        color: #000000 !important;
		        cursor: pointer;
		        display: inline-block;
		        font-size: 15px;
		        font-weight: bold;
		        padding: 5px 10px;
		        text-decoration: none;
		        vertical-align: top; 
		        background-color: #e6e6e6;   
		    }
		    a.btn{font-size: 14px;}
		   
		    .btn:hover{
		        box-shadow: 0 0 3px #888;
		        color: #fff !important; 
		        background-color: #369; 
		        cursor: pointer;   
		    }
		    .btn h2{
		        cursor: pointer;   
		    }
		    .del{
		        background-color: #bbb; 
		    }
		    .link{
		        cursor: pointer;  
		        text-decoration: none;     
		    }
		    .newlogin{
		        cursor: pointer;  
		        font-size: 20px; 
		        background-color: #77b;
		        margin:3px;
		            
		    }     
		    .del:hover{
		        background-color: #f00; 
		    } 
		    .link:hover{
		        color: #f00; 
		    }
		    .promotion{

		        border-radius: 2px;
		        box-shadow: 0 0 3px #333;
		        color: #000000 !important;
		        cursor: pointer;
		        display: inline-block;
		        font-size: 15px;
		        font-weight: bold;
		        padding: 5px 10px;
		        text-decoration: none;
		        vertical-align: top; 
		        background-color: #e6e6e6;   
		    }           

		</style> 
		<div style="float:left; width:20%;text-align:center;"> <h2> Số Tool: 696969 </h2> </div><div style="float:left; width:25%;"> <div> Loại tool: <b>Free </b></b></div> <div>Thời hạn sử dụng: <b>696969 ngày </b> </div>   <div><b>Hỗ trợ: 0926.188.188 </b> (SMS bạn nhé)</div> 
		    <div><a href="/toolv3/recode.php"><b>Gửi Lại Mã Kích Hoạt </b></a></div> 
		     </div> <div style="float:left; text-align:left; width:35%;">
		     
		    <!--Gia hạn bằng thẻ Gate, Viettel, Vina, Mobi... <br /> Thanh toán qua ngân hàng tặng 10% ngày sử dụng.--></div>
		    <!--
		    <div style="float:right; text-align:right; width:15%;">
		    <a class="btn" href="/toolv3/pay.php"> <h2>Gia Hạn</h2> </a> </div> 
		    -->
		    
		    
		    <br style="clear:both; margin:5px;" /> <hr style="clear:both;" /> <center>
		    
		    <h1 style="font-size: 20px; color:red;"> </h1>

		 </center>


		        <h1 style="font-size: 20px; color:blue;">
		        HOT: Chuyển khoản gia hạn ngay với khuyến mãi sau bạn nhé:
		        <br />Tool 100k: <font size="4" color="red">Chuyển 500k = 365 ngày, 200k = 99 ngày, 100k = 36 ngày. </font>
		        <br />Tool 200k và 300k: <font size="4" color="red"> Chuyển 1tr = 365 ngày, 500k = 99 ngày, 200k = 36 ngày. </font>
		        </h1>
		        
		(Nội dung chuyển khoản:  0382283663 Gia han tool 696969)     
		<br /> <a class="btn" onclick="document.getElementById('bank').style.display='block';">
		<b>Xem thông tin chuyển khoản &gt;&gt;</b> </a>
		<br /><div id="bank" style="display:none;"><div style="margin: 20px 5px 10px 0px;border-radius: 2px;box-shadow: 0 0 3px #333;color: #000000 !important;cursor: default;display: inline-block; padding: 20px 20px;text-decoration: none;-align: top; width:920px;font-size:15px;"> <h3 style="margin: 0px 0px 10px 50px">Tài khoản ngân hàng</h3> <h4>1. Zalopay/Momo: 0926188188 </h4><h4>2. Vietcombank chi nhánh TPHCM. STK: 0531000292106. Chủ TK: Nguyễn Thị Ngọc Diệp</h4><h4>3. Agribank chi nhánh TPHCM. STK: 6380205214891. Chủ TK: Nguyễn Thị Ngọc Diệp</h4><h4>4. VietinBank chi nhánh TPHCM. STK: 102006876517. Chủ TK: Nguyễn Thị Ngọc Diệp</h4><h4>5. Sacombank chi nhánh TPHCM. STK: 060102573024. Chủ TK: Nguyễn Thị Ngọc Diệp</h4><!--<h4>6. Đông Á Bank chi nhánh TPHCM. STK: 0102815577. Chủ TK: Nguyễn Đức Anh</h4>   --><br /><br /><font color="red">Khi chuyển tiền, bạn hãy điền <b>Số Điện Thoại của bạn</b> vào phần nội dung chuyển tiền của ngân hàng để chúng tôi tiện liên hệ.</font> <br /><br /></div></div> 



		        
		         <hr style="clear:both;" /> 
		        <div><a class="btn newlogin" href="http://360plus.gn.zing.vn">Gunny 360Game</a>
		        <a class="btn newlogin" href="https://www.mygame.in.th/play?option=boomz">Boomz</a>
		        <a class="btn newlogin" href="http://th.gamagic.com/login/?nexturl=%2FGameLogin%2F%3FGameID%3D9%26ServerID%3D10">DDT Thai</a>
		        <a class="btn newlogin" href="http://play.babygunny.com/">BabyGunny [Private]</a>
		        <a class="btn newlogin" href="http://play.gunnyvo.com/">GunnyVO [Pri]</a>
		        <a class="btn newlogin" href="http://gunnyazo.us/">GunnyAzo [Pri]</a>
		        <a class="btn newlogin" href="http://play.ga2019.com">HomeGunny [Pri]</a>
		        <a class="btn newlogin" href="http://gun321.com/game/">Gun321 [Pri]</a>
		        <a style="float:right;" class="btn newlogin" href="http://gahay.com/">GaHay [Pri]</a>
		        <a style="float:right;" class="btn newlogin" href="http://play.wingunny.com/">WinGunny [Pri]</a>
				<a style="float:right;" class="btn newlogin" href="http://gunnyxua.com/">GunnyXua [Pri]</a>
				<a style="float:right;" class="btn newlogin" href="http://gakyuc.com/">GaKyUc.com [Pri]</a>
				<a style="float:right;" class="btn newlogin" href="http://gunnylau360.net/play/">GunnyLau360 [Pri]</a>
				<a style="float:right;" class="btn newlogin" href="http://play.gunnytop.com/">GunnyTop [Pri]</a>
				<a style="float:right;" class="btn newlogin" href="https://gunnyfun.com">GunnyFun [Pri]</a>


		        </div>

		        <div id="savedacc">
		        	<div style="float:left; width:40%;">
		        		<h3>Vào nhanh các tài khoản đã đăng nhập: </h3>
		        	</div>
		        	<div style="float:left; width:40%;">
		        		</div><br style="clear:both;" />
		`)

        savedFile, _ := os.OpenFile("static\\db\\saved", os.O_RDWR, 0644)
        defer savedFile.Close()
        scanner := bufio.NewScanner(savedFile)
	    for scanner.Scan() {
	        ctx.Writef(`<div id="acc_` + scanner.Text() + `" style="float:left;padding-left:5em">`)
	        ctx.Writef(`<a class="btn del" title="Xóa tài khoản ` + scanner.Text() + ` ra khỏi danh sách" href="delacc://` + scanner.Text() + `@gunny"> X </a>`)
	    	ctx.Writef(`<a class="btn" title="Click để đăng nhập tài khoản ` + scanner.Text() + `" href="goddt://` + scanner.Text() + `@gunny#http://idgunny.360game.vn/login-game"> ` + scanner.Text() + ` </a></div>`)
	    }
       	// footer
    	ctx.Writef(`
       	</div>
        	<iframe style="display:none;" width="0" height="0" src="/toolv3/setup.php"></iframe> </body><iframe style="display:none;" width="0" height="0" src="/toolv3/toolconfig.php"></iframe> 
    

    
		<script type="text/javascript">
		function goServer(server){

		}


		</script>

     `)
    })

    app.Get("/toolv3/setup.php", func(ctx iris.Context) {
        ctx.Writef(`RSA2048cipher-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`)
    })

    app.Get("/toolv3/toolconfig.php", func(ctx iris.Context) {
        ctx.Writef(`pro=1`)
    })

    app.Get("/toolv3/io.php", func(ctx iris.Context) {
        act := ctx.FormValue("act")
        if act == "getconfig" {
        	ctx.Writef(`@dt_line=1@dt_aim=1@dt_x3=1@dt_autowb=1@dt_map=1@dt_hide=1@dt_nguyentieu=0@dt_time=30@dt_angle=360@dt_freeticket=0@dt_bossvip=0@dt_autodtds=0@dt_bossguild=0@dt_saomm=0@dt_tranhba=0@dt_mcmathach=0`)
        } else if act == "getconfig" {
        	ctx.Writef(`MSG#ok#`)
        } else if act == "usersave" {
        	u := ctx.FormValue("u")
        	if u == "" {
        		return
        	}

	        savedFile, _ := os.OpenFile("static\\db\\saved", os.O_RDWR, 0644)
	        defer savedFile.Close()
	        scanner := bufio.NewScanner(savedFile)
		    for scanner.Scan() {
		    	if u == scanner.Text() {
		    		return
		    	}
		    }

    		savedFile.Write([]byte(u + "\n"))
        	ctx.Writef(`MSG#ok#`)
        } else if act == "delacc" {
        	u := ctx.FormValue("id")
        	if u == "" {
        		return
        	}

        	savedFile, _ := os.OpenFile("static\\db\\saved", os.O_RDWR, 0644)
	        defer savedFile.Close()
	        scanner := bufio.NewScanner(savedFile)
	        line := -1
		    for scanner.Scan() {
		    	line += 1
		    	if u == scanner.Text() {
		    		removeLine("static\\db\\saved", line)
		    		return
		    	}
		    }
		    return

        } else if act == "getdameboss"{
        	savedFile, _ := os.OpenFile("static\\db\\dameboss", os.O_RDWR, 0644)
	        defer savedFile.Close()
	        scanner := bufio.NewScanner(savedFile)
		    for scanner.Scan() {
		    	ctx.Writef(scanner.Text())
		    }
        } else {
        	ctx.Writef(`MSG#ok#`)
        }
    })
    // app.Get("/toolv3/home.php", func(ctx iris.Context) {
    //     ctx.Writef(``)
    // })

    app.Post("/toolv3/ocr.php", func(ctx iris.Context) {
        ctx.Writef(`<winds><wind>0</wind><wind>1</wind><wind>2</wind><wind>3</wind><wind>4</wind><wind>5</wind><wind>6</wind><wind>7</wind><wind>8</wind><wind>9</wind><wind>.</wind></winds>`)
    })

    app.Get("/toolv3/damebossconfig.php", func(ctx iris.Context) {
    	id := ctx.FormValue("id")
    	del := ctx.FormValue("del")
    	if del != "" {
    		config := ""
    		newconfig := ""
		    savedFile, _ := os.OpenFile("static\\db\\dameboss", os.O_RDWR, 0644)
	        scanner := bufio.NewScanner(savedFile)
		    for scanner.Scan() {
		    	config = scanner.Text()
		    	c1 := strings.Split(config, "#")

		    	for _, c2 := range c1 {
		    		if c2 == "" {
		    			continue
		    		}
		    		c3 := strings.Split(c2, "@")
		    		if del == c3[0] {
		    			//found
		    			continue
		    		}

		    		newconfig = newconfig + c2 + "#"
		    	}
		    }
		    savedFile.Close()
		    os.Remove("static\\db\\dameboss")
		    newFile, err := os.OpenFile("static\\db\\dameboss", os.O_CREATE, 0644)
		    if err != nil {
		    	fmt.Println("err: " + err.Error())
		    }
        	defer newFile.Close()
        	newFile.Write([]byte(newconfig))
		    ctx.Writef(`<script>window.location='/toolv3/damebossconfig.php'</script>`)
    	}
    	if id ==  "" {
    		ctx.Writef(`<body style="margin:10px;background-color: #eee;">
        	<style type="text/css">
			    * {
			        cursor: default;
			        -webkit-user-select: none;     
			    }  
			    .editform input    { -webkit-user-select: auto !important;}
			    .editform  { margin: 5px 100px;}
			    .btn{

			        border-radius: 2px;
			        box-shadow: 0 0 3px #333;
			        color: #000000 !important;
			        cursor: pointer;
			        display: inline-block;
			        font-size: 15px;
			        font-weight: bold;
			        padding: 5px 10px;
			        text-decoration: none;
			        vertical-align: top; 
			        background-color: #e6e6e6;   
			    }
			   
			    .btn:hover{
			        box-shadow: 0 0 3px #888;
			        color: #fff !important; 
			        background-color: #369; 
			        cursor: pointer;   
			    }

			    .link:hover{
			        color: #f00; 
			    } 
			    .del{text-decoration: none; color:#c00;  }
			    input {text-transform: uppercase;} 
		    </style> 
		    <a title="tạo mới" href="/toolv3/damebossconfig.php?id=999999"> <h2>Tạo mới</h2></a>
		    <h2> Các tùy chỉnh buff dame boss đã lưu</h2>
		    `)
		config := ""
	    savedFile, _ := os.OpenFile("static\\db\\dameboss", os.O_RDWR, 0644)
        defer savedFile.Close()
        scanner := bufio.NewScanner(savedFile)
	    for scanner.Scan() {
	    	config = scanner.Text()
	    	c1 := strings.Split(config, "#")

	    	for _, c2 := range c1 {
	    		if c2 == "" {
	    			continue
	    		}
	    		c3 := strings.Split(c2, "@")
	    		ctx.Writef(`<div style="float:left;padding-left:5em"> 
			    	<a title="Xóa" class="del" onclick="if(confirm('Chắc chắn xóa tùy chỉnh này?')){return true;}else{return false;}" href="/toolv3/damebossconfig.php?del=` + c3[0] +`">x</a> 
			    	<a class="btn" title="` + c3[0] +`" href="/toolv3/damebossconfig.php?id=` + c3[0] +`">` + c3[0] +`</a>
		    	</div>`)
	    		for index, _ := range c3 {
	    			if index == 0 {

	    			}
	    		}
	    	}
	    }

		ctx.Writef(`
		    </div></div></body>`)
    	} else if id == "999999"{
    		ctx.Writef(`<body style="margin:10px;background-color: #eee;"><style type="text/css">
			    * {
			        cursor: default;
			        -webkit-user-select: none;     
			    }  
			    .editform input    { -webkit-user-select: auto !important;}
			    .editform  { margin: 5px 100px;}
			    .btn{

			        border-radius: 2px;
			        box-shadow: 0 0 3px #333;
			        color: #000000 !important;
			        cursor: pointer;
			        display: inline-block;
			        font-size: 15px;
			        font-weight: bold;
			        padding: 5px 10px;
			        text-decoration: none;
			        vertical-align: top; 
			        background-color: #e6e6e6;   
			    }
			   
			    .btn:hover{
			        box-shadow: 0 0 3px #888;
			        color: #fff !important; 
			        background-color: #369; 
			        cursor: pointer;   
			    }

			    .link:hover{
			        color: #f00; 
			    } 
			    .del{text-decoration: none; color:#c00;  }
			    input {text-transform: uppercase;} </style> <div style="font-size:25px;"> Tùy chỉnh dame boss </div><form method="post" action="?id=999999" class="editform">
			        <b>Đặt tên cho kiểu buff</b>: <input type="text" name="name" value="Dame Boss 999999" size="38" maxlength="20" /><br />
			        <b>Buff Turn 1</b> <input maxlength="20" type="text" name="turn1" value="23444445678" size="50" /><br />
			        <b>Buff Turn 2</b> <input maxlength="20" type="text" name="turn2" value="23444445678" size="50" /><br />
			        <b>Buff Turn 3</b> <input  maxlength="20" type="text" name="turn3" value="Q3444445678" size="50" /><br />
			        <b>Buff Turn 4</b> <input  maxlength="20" type="text" name="turn4" value="B3444445678" size="50" /><br />
			        <b>Buff Turn 5</b> <input  maxlength="20" type="text" name="turn5" value="23444445678" size="50" /><br />
			        <b>Buff Turn 6</b> <input  maxlength="20" type="text" name="turn6" value="23444445678" size="50" /> <br />
			        <b>Buff Turn 7</b> <input  maxlength="20" type="text" name="turn7" value="23444445678" size="50" /> <br />
			        <b>Buff Turn 8</b> <input  maxlength="20" type="text" name="turn8" value="23444445678" size="50" /> <br />
			        <b>Buff Turn 9</b> <input  maxlength="20" type="text" name="turn9" value="23444445678" size="50" /> <br />
			        <b>Buff Turn 10</b> <input  maxlength="20" type="text" name="turn10" value="23444445678" size="50" /> <br />
			        <b>Buff Turn 11</b> <input  maxlength="20" type="text" name="turn11" value="23444445678" size="50" /> <br />
			        <b>Buff Turn 12</b> <input  maxlength="20" type="text" name="turn12" value="23444445678" size="50" /> <br />
			        <b>Buff Turn 13</b> <input  maxlength="20" type="text" name="turn13" value="23444445678" size="50" /> <br />
			        <b>Buff Turn 14</b> <input  maxlength="20" type="text" name="turn14" value="23444445678" size="50" /> <br />
			        <b>Buff Turn 15</b> <input  maxlength="20" type="text" name="turn15" value="23444445678" size="50" /> <br />
			        <b>Buff Turn 16</b> <input  maxlength="20" type="text" name="turn16" value="23444445678" size="50" /> <br />
			        <b>Buff Turn 17</b> <input  maxlength="20" type="text" name="turn17" value="23444445678" size="50" /> <br />
			        <b>Buff Turn 18</b> <input  maxlength="20" type="text" name="turn18" value="23444445678" size="50" /> <br />
			        <b>Buff Turn 19</b> <input  maxlength="20" type="text" name="turn19" value="23444445678" size="50" /> <br />
			        <b>Buff Turn 20</b> <input  maxlength="20" type="text" name="turn20" value="23444445678" size="50" /> <br />
			        <b>Buff Turn 21</b> <input  maxlength="20" type="text" name="turn21" value="23444445678" size="50" /> <br />
			        <b>Buff Turn 22</b> <input  maxlength="20" type="text" name="turn22" value="23444445678" size="50" /> <br />
			        <b>Buff Turn 23</b> <input  maxlength="20" type="text" name="turn23" value="23444445678" size="50" /> <br />
			        <b>Buff Turn 24</b> <input  maxlength="20" type="text" name="turn24" value="23444445678" size="50" /> <br />
			        <b>Buff Turn 25</b> <input  maxlength="20" type="text" name="turn25" value="23444445678" size="50" /> <br />
			        
			        
			        <input type="hidden" name="id_post" value="999999" />
			        <input type="submit" value="Lưu" />  <=====> <a href="/toolv3/damebossconfig.php">Trở về danh sách</a> 
			        <br />

			        Ghi chú: Để buff <b>Pow</b> dùng phím <b>B</b>, Skill pet dùng các phím <b>Q-W-E-R-T</b> 
			        <br /> Buff kích pow hoặc trọng thương dùng các phím <b>Z-X-C</b>
			        <br /> Thứ tự ưu tiên từ trái sang phải (buff chiêu từ trái sang phải cho tới hết thể lực)
			        </form>      
			        </div></body>`)
    	} else {
    		ctx.Writef(`<body style="margin:10px;background-color: #eee;"><style type="text/css">
			    * {
			        cursor: default;
			        -webkit-user-select: none;     
			    }  
			    .editform input    { -webkit-user-select: auto !important;}
			    .editform  { margin: 5px 100px;}
			    .btn{

			        border-radius: 2px;
			        box-shadow: 0 0 3px #333;
			        color: #000000 !important;
			        cursor: pointer;
			        display: inline-block;
			        font-size: 15px;
			        font-weight: bold;
			        padding: 5px 10px;
			        text-decoration: none;
			        vertical-align: top; 
			        background-color: #e6e6e6;   
			    }
			   
			    .btn:hover{
			        box-shadow: 0 0 3px #888;
			        color: #fff !important; 
			        background-color: #369; 
			        cursor: pointer;   
			    }

			    .link:hover{
			        color: #f00; 
			    } 
			    .del{text-decoration: none; color:#c00;  }
			    `)

			config := ""
		    savedFile, _ := os.OpenFile("static\\db\\dameboss", os.O_RDWR, 0644)
	        defer savedFile.Close()
	        scanner := bufio.NewScanner(savedFile)
		    for scanner.Scan() {
		    	config = scanner.Text()
		    	c1 := strings.Split(config, "#")

		    	for _, c2 := range c1 {
		    		if c2 == "" {
		    			continue
		    		}
		    		c3 := strings.Split(c2, "@")
		    		if c3[0] == id {
		    			ctx.Writef(`input {text-transform: uppercase;} </style> <div style="font-size:25px;"> Tùy chỉnh dame boss </div><form method="post" action="?id=` + c3[0] + `" class="editform">
		    				<b>Đặt tên cho kiểu buff</b>: <input type="text" name="name" value="` + c3[0] + `" size="38" maxlength="20" /><br />`)
		    			for i := 1; i < 26; i++ {
		    				ctx.Writef(`<b>Buff Turn ` + strconv.Itoa(i) + `</b> <input maxlength="20" type="text" name="turn` + strconv.Itoa(i) + `" value="` + c3[i] + `" size="50" /><br />`)
		    			} 
		    			ctx.Writef(`<input type="hidden" name="id_post" value="` + c3[0] + `" />`)
		    		}
		    	}
		    }
		    ctx.Writef(`
		        <input type="submit" value="Lưu" />  <=====> <a href="/toolv3/damebossconfig.php">Trở về danh sách</a> 
		        <br />

		        Ghi chú: Để buff <b>Pow</b> dùng phím <b>B</b>, Skill pet dùng các phím <b>Q-W-E-R-T</b> 
		        <br /> Buff kích pow hoặc trọng thương dùng các phím <b>Z-X-C</b>
		        <br /> Thứ tự ưu tiên từ trái sang phải (buff chiêu từ trái sang phải cho tới hết thể lực)
		        </form>      
		        </div></body>`)
    	}
    })


    app.Post("/toolv3/damebossconfig.php", func(ctx iris.Context) {
    	id := ctx.URLParam("id")
    	if id == "" {
    		return
    	}
    	if id == "999999" {
    		config := ctx.FormValue("name")
    		for i := 1; i < 26; i++ {
    			config += "@"
    			config += ctx.FormValue("turn" + strconv.Itoa(i))
    		} 
    		fmt.Println(config)
    		damebossconfig, _ := os.OpenFile("static\\db\\dameboss", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0720)
    		defer damebossconfig.Close()
    		damebossconfig.Write([]byte("#" + config))
    		ctx.Writef(`<script>window.location='/toolv3/damebossconfig.php'</script>`)
    		return
    	}
    	//else
    	config := ""
		newconfig := ""
	    savedFile, _ := os.OpenFile("static\\db\\dameboss", os.O_RDWR, 0644)
        scanner := bufio.NewScanner(savedFile)
	    for scanner.Scan() {
	    	config = scanner.Text()
	    	c1 := strings.Split(config, "#")

	    	for _, c2 := range c1 {
	    		if c2 == "" {
	    			continue
	    		}
	    		c3 := strings.Split(c2, "@")
	    		if ctx.FormValue("name") == c3[0] {
	    			//found
	    			newconfig += c3[0]
	    			for i := 1; i < 26; i++ {
	    				newconfig += "@"
		    			newconfig += ctx.FormValue("turn" + strconv.Itoa(i))
		    		}
	    			newconfig += "#"
	    			continue
	    		}

	    		newconfig = newconfig + c2 + "#"
	    	}
	    }
	    savedFile.Close()
	    os.Remove("static\\db\\dameboss")
	    newFile, err := os.OpenFile("static\\db\\dameboss", os.O_CREATE, 0644)
	    if err != nil {
	    	fmt.Println("err: " + err.Error())
	    }
    	defer newFile.Close()
    	newFile.Write([]byte(newconfig))
	    ctx.Writef(`<script>window.location='/toolv3/damebossconfig.php'</script>`)




    })

    app.Run(iris.TLS("127.0.0.1:443", "static\\service.crt", "static\\service.key"))
}


func removeLine(path string, lineNumber int) {
    file, err := ioutil.ReadFile(path)
    if err != nil {
        panic(err)
    }

    info, _ := os.Stat(path)
    mode := info.Mode()

    array := strings.Split(string(file), "\n")
    array = append(array[:lineNumber], array[lineNumber+1:]...)
    ioutil.WriteFile(path, []byte(strings.Join(array, "\n")), mode)
}
