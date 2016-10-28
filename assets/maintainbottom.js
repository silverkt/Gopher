// 调试功能    
$("#debug").click(function(){
    $sc = $("#myiframe").prop("src"); 
    $str = "open the debug tool of "+$sc;
     $inde = $('#proj li[class="active"]');
    $ind = $('#proj li').index($inde);
    if($ind == 0) {$out = "腾讯上海云数据中心";}
    if($ind == 1) {$out = "肇庆新区泛能站";}
    if($ind == 2) {$out = "黄花机场泛能站";}
    if($ind == 3) {$out = "亭湖医院泛能站";}
    if($ind == 4) {$out = "株洲神农城泛能站";}
    if($ind == 5) {$out = "中德2#泛能站";}
    if($ind == 6) {$out = "株洲职教城泛能站";}
    if($ind == 7) {$out = "振西商贸广场泛能站";}
    if($ind == 8) {$out = "海宁航海国际项目";}
    if($ind == 9) {$out = "新朝阳泛能微网";}
    if($ind == 10) {$out = "蚌埠大明产业园";}    
            
              
    $str = "打开调试 "+$out;
    alert($str);
});
         
 