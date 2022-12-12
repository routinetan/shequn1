var swiper1 = new Swiper('.swiper1', {
    slidesPerView: 3.8,
    loop: false,//是否可循环
});
setTimeout(function () {
    mpContent();
}, 300);

setTimeout(function () {
    navContent();
}, 350);
setTimeout(function () {
    qunSumContent();
}, 400);
setTimeout(function () {
    ul1Content();
}, 450);

function mpContent() {    //公众号的内容
    var mp = '';
    mp += '<div class="bcardz"  style="background-color:#666666;"  >';
    mp += '	<div class="bcard">';
    mp += '		<div class="llff">';
    mp += '			<img src="https://www.haobangyingxiaocehua.cn/attachment/images/2/2020/10/cquUVq601Ud0Knb6YvBkvzUxLunB11.jpg"/>';
    mp += '		</div>';
    mp += '		<div class="lcc">';
    mp += '			<p  style="color:#ffffff;" >好帮营销策划</p>';
    mp += '			<p   style="color:#ff0000;"  class="lct">企业线上拓客营销策划等服务！</p>';
    mp += '		</div>';
    mp += '		<div class="lr">';
    mp += '			<div class="popup" id="pop" style="top:0;z-index:5;">';
    mp += '				<div class="popc" style="margin:35% auto;">';
    mp += '					<img src="/static/image/delet.png" class="delet" id="closebutt" onclick="closeButton(1)"/>';
    mp += '					<span>好帮营销策划</span>';
    mp += '					<img src="http://haobang81.oss-cn-beijing.aliyuncs.com/images/2/2020/08/GZ63SwwGUgsiSQJ1ouo7xP7x31QWss.jpg" class="eq"/>';
    mp += '					<p>长按二维码</p>';
    mp += '				</div>';
    mp += '			</div>';
    mp += '			<input type="button" style=" background-color:#ffffff;  color:#000000;"  value="关注公众号" id="openbutt" onclick="clickButton()" class="inp"/>';
    mp += '		</div>';
    mp += '	</div>';
    mp += '</div>';
    $("#mp").append(mp);
}

function navContent() {    //分类导航
    var nav = '';
    nav += '<div class="subnav">';
    nav += '	<ul>';
    nav += '		<a href="https://www.haobangyingxiaocehua.cn/">';
    nav += '			<li>';
    nav += '				<img src="https://www.haobangyingxiaocehua.cn/attachment/images/2/2022/03/vc69H9ZRY6WrPrAc7rpUlZa76HRr66.jpg"/>';
    nav += '				<p>好帮营销</p>';
    nav += '			</li>';
    nav += '		</a>';
    nav += '		<a href="http://www.haobangyingxiaocehua.cn/cmd/index.php?i=2&controller=entry&act=index&do=user.sys_case&m=lywywl_ztb">';
    nav += '			<li>';
    nav += '				<img src="https://www.haobangyingxiaocehua.cn/attachment/images/2/2020/12/Bw4uU1OszFS6ZAsOXuS1C1nU6wsWEF.jpg"/>';
    nav += '				<p>营销工具</p>';
    nav += '			</li>';
    nav += '		</a>';
    nav += '		<a href="https://www.haobangyingxiaocehua.cn/cmd/index.php?i=2&controller=entry&do=index&m=nan_fc&state=1">';
    nav += '			<li>';
    nav += '				<img src="https://www.haobangyingxiaocehua.cn/attachment/images/2/2020/12/wwzZgUk6w6YsGwyZuEu96oY6UUY76y.jpg"/>';
    nav += '				<p>娱乐游戏</p>';
    nav += '			</li>';
    nav += '		</a>';
    nav += '		<a href="https://www.haobangyingxiaocehua.cn/./index.php?i=2&controller=entry&do=index&m=small_poster&id=19">';
    nav += '			<li>';
    nav += '				<img src="https://www.haobangyingxiaocehua.cn/attachment/images/2/2020/12/LfpCiSuimO85ZwSc1C5SIJWQUsUSIF.jpg"/>';
    nav += '				<p>微信商城</p>';
    nav += '			</li>';
    nav += '		</a>';
    nav += '		<a href="https://www.haobangyingxiaocehua.cn/./index.php?i=2&controller=entry&do=index&m=small_poster&id=18">';
    nav += '			<li>';
    nav += '				<img src="https://www.haobangyingxiaocehua.cn/attachment/images/2/2020/12/MoZj6BN2m6Xd40gmBE22gomX4j0D66.jpg"/>';
    nav += '				<p>客服微信</p>';
    nav += '			</li>';
    nav += '		</a>';
    nav += '	</ul>';
    nav += '</div>';
    $("#nav").append(nav);
}

function qunSumContent() {    //群统计
    var qunSum = '';
    qunSum += '<div class="buyqun">';
    qunSum += '	<div class="quncont">本站共有<span class="qunsum">1980</span>个 本地群 ，196586个群成员</div>';
    qunSum += '	<div class="buyvip"><a href="./index.php?i=2&controller=entry&do=vipContent&m=yy_shequn2"> 点击购买VIP,拉您进20个社群并尊享会员特权 &gt</a></div>';
    qunSum += '	<img src="/static/image/vip1.png" class="buybg"/>';
    qunSum += '</div>';
    $("#qunSum").append(qunSum);
}
function ul1Content() {    //初始群数据
    var urlLink = "/group_list";
    $.ajax({
        url: urlLink,
        type: "get",
        dataType: 'json',
        success: function (ret) {
            data = ret.data
            if (ret.rows > 0) {
                for (var j = 0; j < data.length; j++) {
                    var ruqunname = getruqunname(data[j].jumpbtntext, "", "入群");
                    var title = data[j].title;
                    var ico = data[j].ico;
                    var label = data[j].label;
                    var itemId = data[j].id;
                    var icoId = itemId % 10;
                    var qunIco = 'qun' + icoId + '.jpg';//根据群ID的末位数来获取对应的群头像文件名
                    var linkUrl = data[j].linkUrl;//外链
                    var type = data[j].type;//社群模式
                    var btnColor = data[j].btnColor;
                    var hintName = data[j].hintName;
                    var hintNote = data[j].hintNote;
                    var hintImg = data[j].hintImg;
                    var hintBtnTitle = data[j].hintBtnTitle;
                    var hintBtnColor = data[j].hintBtnColor;
                    var hintImgLaber = data[j].hintImgLaber;
                    var hint_type = data[j].hint_type;
                    var hintImgUrl = "http://haobang81.oss-cn-beijing.aliyuncs.com/" + hintImg;
                    var aLink = "/group/info?itemId=" + itemId; //群详情页
                    var html2 = '';
                    html2 += '<li>';
                    html2 += '<div class="popup" id="popup' + itemId + '">';
                    html2 += '	<div class="popc" style="margin:25% auto;">';
                    html2 += '		<img src="/static/image/delet.png" class="delet"  onclick="closeBut(' + itemId + ')"/>';
                    if (hint_type == 1) {  //调用自己弹窗
                        html2 += '<span>' + hintName + '</span>';
                        html2 += '<img src="' + hintImgUrl + '" class="eq"/>';
                        html2 += '<p>' + hintImgLaber + '</p>';
                        html2 += '<p>' + hintNote + '</p>';
                        html2 += '</div>';
                        html2 += '<div class="joinqz" style="width:100%;margin:0 auto;">';
                        if (hintBtnTitle != null && hintBtnTitle != undefined && hintBtnTitle != '') {
                            html2 += '<a href=' + aLink + ' class="joinq" style="background-color:#12a72e;">' + hintBtnTitle + '></a>';
                        }

                    } else if (hint_type == 0) {  //跟随调用系统弹窗
                        html2 += '<span>模具体雕美体内衣</span>';
                        html2 += '<img src="http://haobang81.oss-cn-beijing.aliyuncs.com/images/2/2020/08/NFhfH60pDfJEEFS0srnhpSqQPRhdR6.jpg" class="eq"/>';
                        html2 += '<p>赠送好礼</p>';
                        html2 += '<span>加我微信182，厂家直销美体内衣模具体雕等！</span>';
                        html2 += '</div>';
                        html2 += '<div class="joinqz" style="width:100%;margin:0 auto;">'
                        html2 += '<a href=' + aLink + ' class="joinq" style="background-color:#12a72e;">不了,我要直接入群></a>';

                    }
                    html2 += '</div>';
                    html2 += '</div>';
                    html2 += '	<div class="llf">';
                    //html2+='		<img src="../addons/yy_shequn2/images/qunIco/'+qunIco+'"/>';
                    /** **/
                    if (data[j].isshequnico == 1 && data[j].shequnico != "") {
                        html2 += '		<img src="' + tomedia_js(data[j].shequnico) + '"/>';
                    }
                    else {
                        html2 += '		<img src="/static/image/' + qunIco + '"/>';
                    }
                    /** **/
                    html2 += '	</div>';
                    html2 += '	<div class="lc">';
                    html2 += '		<p class="lcp">' + title + '</p>';
                    html2 += '		<p class="lct">' + label + '</p>';
                    html2 += '	</div>';
                    if (btnColor != null && btnColor != undefined && btnColor != '') {
                        html2 += '	<div  style="color: ' + btnColor + ';border: 0.15rem solid ' + btnColor + ';width: 23%;text-align: center;width: 3.75rem;height: 1.82rem;border-radius: 1rem;font-size: 0.857rem;font-weight: bold;"  class="lr" >';
                    } else {
                        html2 += '	<div  style="color: #1aaf78;border: 0.15rem solid #1aaf78;width: 23%;text-align: center;width: 3.75rem;height: 1.82rem;border-radius: 1rem;font-size: 0.857rem;font-weight: bold;"   class="lr" >';
                    }
                    if (hint_type == 0) {  //跟随社群统一设置类型
                        html2 += '<a class="ruqunBtnLink" onclick="autocpy()" href=' + aLink + '>' + ruqunname + '</a>';
                    } else if (hint_type == 1) {//调用自己弹窗
                        html2 += '<div  onclick="autocpy();clickBut(' + itemId + ')" class="ruqunBtn">' + ruqunname + '</div>';
                    } else if (hint_type == 2) {//直接入群
                        html2 += '<a class="ruqunBtnLink" onclick="autocpy()" href=' + aLink + '>' + ruqunname + '</a>';
                    } else if (hint_type == 3) {//外链
                        html2 += '<a class="ruqunBtnLink" onclick="autocpy()" href=' + linkUrl + '>' + ruqunname + '</a>';
                    }

                    html2 += '	</div>';
                    html2 += '</li>';
                    $(html2).appendTo($("#ul1"));

                }
            }
        },
        error: function () {
            console.log("出错了");
        }
    });
}


var page = 2;
var num = 8;
hui.loadMore(getMore);
function getMore() {//上拉到底时触发
    hui.get(
        "./group_list?page=" + page+"&num="+num,
        function (res) {
            //判断加载完毕

            var ret = JSON.parse(res);
            if (ret.rows == 0) {
                hui.endLoadMore(true, '.');
                console.log('lllllli23333iiiiihhhhhhhh');
                $('.pull-loading').html("没有了哟");
                return false;
            }
            data = ret.data
            for (var j = 0; j < data.length; j++) {
                var ruqunname = getruqunname(data[j].jumpbtntext, "", "入群");
                var title = data[j].title;
                var ico = data[j].ico;
                var label = data[j].label;
                var itemId = data[j].id;
                var icoId = itemId % 10;
                var qunIco = 'qun' + icoId + '.jpg';//根据群ID的末位数来获取对应的群头像文件名
                var linkUrl = data[j].linkUrl;//外链
                var type = data[j].type;//社群模式
                var btnColor = data[j].btnColor;
                var hintName = data[j].hintName;
                var hintNote = data[j].hintNote;
                var hintImg = data[j].hintImg;
                var hintBtnTitle = data[j].hintBtnTitle;
                var hintBtnColor = data[j].hintBtnColor;
                var hintImgLaber = data[j].hintImgLaber;
                var hint_type = data[j].hint_type;
                var hintImgUrl = "http://haobang81.oss-cn-beijing.aliyuncs.com/" + hintImg;
                var aLink = "/group/info?itemId=" + itemId; //群详情页
                var html2 = '';
                html2 += '<li>';
                html2 += '<div class="popup" id="popup' + itemId + '">';
                html2 += '	<div class="popc" style="margin:25% auto;">';
                html2 += '		<img src="/static/image/delet.png" class="delet"  onclick="closeBut(' + itemId + ')"/>';
                if (hint_type == 1) {  //调用自己弹窗
                    html2 += '<span>' + hintName + '</span>';
                    html2 += '<img src="' + hintImgUrl + '" class="eq"/>';
                    html2 += '<p>' + hintImgLaber + '</p>';
                    html2 += '<p>' + hintNote + '</p>';
                    html2 += '</div>';
                    html2 += '<div class="joinqz" style="width:100%;margin:0 auto;">';
                    if (hintBtnTitle != null && hintBtnTitle != undefined && hintBtnTitle != '') {
                        html2 += '<a href=' + aLink + ' class="joinq" style="background-color:#12a72e;">' + hintBtnTitle + '></a>';
                    }
                } else if (hint_type == 0) {  //跟随调用系统弹窗
                    html2 += '<span>模具体雕美体内衣</span>';
                    html2 += '<img src="http://haobang81.oss-cn-beijing.aliyuncs.com/images/2/2020/08/NFhfH60pDfJEEFS0srnhpSqQPRhdR6.jpg" class="eq"/>';
                    html2 += '<p>赠送好礼</p>';
                    html2 += '<span>加我微信182，厂家直销美体内衣模具体雕等！</span>';
                    html2 += '</div>';
                    html2 += '<div class="joinqz" style="width:100%;margin:0 auto;">';
                    html2 += '<a href=' + aLink + ' class="joinq" style="background-color:#12a72e;">不了,我要直接入群></a>';
                }
                html2 += '     </div>';
                html2 += '</div>';
                html2 += '	<div class="llf">';
                //html2+='		<img src="../addons/yy_shequn2/images/qunIco/'+qunIco+'"/>';
                /** **/
                if (data[j].isshequnico == 1 && data[j].shequnico != "") {
                    html2 += '		<img src="' + tomedia_js(data[j].shequnico) + '"/>';
                }
                else {
                    html2 += '		<img src="/static/image/' + qunIco + '"/>';
                }
                /** **/
                html2 += '	</div>';
                html2 += '	<div class="lc">';
                html2 += '		<p class="lcp">' + title + '</p>';
                html2 += '		<p class="lct">' + label + '</p>';
                html2 += '	</div>';
                if (btnColor != null && btnColor != undefined && btnColor != '') {
                    html2 += '	<div  style="color: ' + btnColor + ';border: 0.15rem solid ' + btnColor + ';width: 23%;text-align: center;width: 3.75rem;height: 1.82rem;border-radius: 1rem;font-size: 0.857rem;font-weight: bold;"  class="lr" >';
                } else {
                    html2 += '	<div  style="color: #1aaf78;border: 0.15rem solid #1aaf78;width: 23%;text-align: center;width: 3.75rem;height: 1.82rem;border-radius: 1rem;font-size: 0.857rem;font-weight: bold;"   class="lr" >';
                }
                if (hint_type == 0) {  //跟随社群统一设置类型
                    html2 += '<a class="ruqunBtnLink" onclick="autocpy();", href=' + aLink + '>' + ruqunname + '</a>';
                } else if (hint_type == 1) {//调用自己弹窗
                    html2 += '<div  onclick="autocpy();clickBut(' + itemId + ')" class="ruqunBtn">' + ruqunname + '</div>';
                } else if (hint_type == 2) {//直接入群
                    html2 += '<a class="ruqunBtnLink" onclick="autocpy();", href=' + aLink + '>' + ruqunname + '</a>';
                } else if (hint_type == 3) {//外链
                    html2 += '<a class="ruqunBtnLink" onclick="autocpy();", href=' + linkUrl + '>' + ruqunname + '</a>';
                }

                html2 += '	</div>';
                html2 += '</li>';
                $(html2).appendTo($("#ul2"));

            }
            page += 1;
            num  +=8;
            hui.endLoadMore();

        }
    );



    // myscroll.refresh();
}

function clickButton() {
    var popupName = 'pop';
    var popup = document.getElementById(popupName);
    popup.style.display = "block";
}
function closeButton() {
    var popupName = 'pop';
    var popup = document.getElementById(popupName);
    popup.style.display = "none";
}

function clickButton(){
    var popupName = 'pop';
    var popup = document.getElementById(popupName);
    popup.style.display="block";
}
function closeButton(){
    var popupName = 'pop';
    var popup=document.getElementById(popupName);
    popup.style.display="none";
}

function getruqunname(fname,cname,dname){
    var v=dname;
    if(fname!=undefined && fname!="")
        v=fname;
    else if(cname!=undefined && cname!="")
        v=cname;
    return v;
}
function tomedia_js(imgurl){
    var url = "http://haobang81.oss-cn-beijing.aliyuncs.com/";
    return url + imgurl;
}
function clickBut(itemId){
    var popupName = 'popup'+itemId;
    var popup = document.getElementById(popupName);
    popup.style.display="block";
}
function closeBut(itemId){
    var popupName = 'popup'+itemId;
    var popup=document.getElementById(popupName);
    popup.style.display="none";
}
function autocpy(){
}
