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
