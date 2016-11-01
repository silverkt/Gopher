
          //动态更改页面中iframe大小
function changeFrameHeight() {
    var ifm = document.getElementById("myiframe"); 
    ifm.height=document.documentElement.clientHeight-210;
    
}
         
window.onresize=function(){  
    
     changeFrameHeight();  
} 


        //动态更改遮罩层iframe大小
function changenew(){
    var ifm= document.getElementById("newiframe"); 
    ifm.height=document.documentElement.clientHeight;
    ifm.width=document.documentElement.clientWidth;
}
window.onresize=function(){  
     changenew();  
} 






        //监听F11事件，动态调整页面大小，避免出现滚动条
        //监听esc事件，关闭遮罩层
        window.onkeyup = function(event){
           
            var e = event || window.event || arguments.callee.caller.arguments[0];
            if(e && e.keyCode == 122) {
                
                changeFrameHeight();  
                changenew(); 
            }
             if(e && e.keyCode == 27) {
               
               hideMask();
            }
        }
        
        
function esciframe(){
           setInterval("getblur()",1000);
        }





    //兼容火狐、IE8 
    //显示遮罩层  
    function showMask(){   
        
        $sc = $("#myiframe").prop("src"); 
        var $jsn = {"src":$sc};
        $("#newiframe").prop($jsn); 
         changenew();  
        $("#mask").css("height",$(document).height());   
        $("#mask").css("width",$(document).width());   
        $("#mask").show(); 
        //设置每隔1秒让全屏遮罩iframe失去焦点，避免onkeyup事件被iframe阻断
        window.setInterval("getblur()",1000);
       
        
    }
    //隐藏遮罩层
    function hideMask(){   
        
        $("#mask").hide();   
    }
    //iframe失去焦点    
    function getblur(){
       $("#newiframe").blur();
       document.getElementById("getfocus").focus();
        
         
    }
   

