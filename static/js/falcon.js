$(document).ready(function() {


    function format(state) {
        return state.id+"<span class='select-info'>("+state.text+")</span>";
    }
    function timeConverter(UNIX_timestamp){

        if (UNIX_timestamp==0) {
            return 0
        }
        var date = new Date(UNIX_timestamp * 1000);

        Y = date.getFullYear() + '-';
        M = (date.getMonth()+1 < 10 ? '0'+(date.getMonth()+1) : date.getMonth()+1) + '-';
        D = date.getDate() + ' ';
        h = date.getHours() + ':';
        m = date.getMinutes() + ':';
        s = date.getSeconds();
        time=Y+M+D+h+m+s;
        return time;    
    }

    $('.falcontpl-relation').click(function(e){
        $("#bindparent_tpl_div").show('fast');
        $("#add_div").hide('fast');
        $("#update").hide('fast');
        $("#uic_update").hide('fast');
        
        var tr = $(e.target).closest('tr');
        var Id = tr.find('.Id').val();
        console.log(Id);
        $("#tpl_id").val(Id);
        
    })

    $('.falcontpl-warm').click(function(e){
        $("#uic_update").show('fast');
        $("#add_div").hide('fast');
        $("#bindparent_tpl_div").hide('fast');
        $("#update").hide('fast');
        
        var tr = $(e.target).closest('tr');
        var Id = tr.find('.Id').val();
        console.log(Id);

        $("#tpl_id").val(Id);

        $.ajax({
                url : '/zjz_work_cmdb/cmdb/falcongetactiondetail',
                method : 'post',
                data: {tpl_id:Id},
                success : function(response){
                    console.log(response)
                    
                    if(response.message==0) {
                        console.log(response.message)
                        $("#action_id").val(response.message);
                        $('#uic_team').html("");
                        $("#callback_url").val("");
                        $("#after_callback_mail").attr("checked",false);
                        $("#after_callback_sms").attr("checked",false);
                        $("#before_callback_mail").attr("checked",false);
                        $("#before_callback_sms").attr("checked",false);
                    } else {
                        $("#after_callback_mail").attr("checked",false);
                        $("#after_callback_sms").attr("checked",false);
                        $("#before_callback_mail").attr("checked",false);
                        $("#before_callback_sms").attr("checked",false);
                       
                        $("#action_id").val(response.message[0].Id);
                        $('#uic_team').html(response.message[0].Uic);
                        
                
                        if (response.message[0].Callback==1){
                            $("#callback_url").val(response.message[0].Url)
                        }
                        if (response.message[0].After_callback_mail==1) {
                            $("#after_callback_mail").prop("checked",true);
                        }
                        if (response.message[0].After_callback_sms==1) {
                            $("#after_callback_sms").prop("checked",true);
                        }
                        if (response.message[0].Before_callback_mail==1) {
                            $("#before_callback_mail").prop("checked",true);
                        }
                        if (response.message[0].Before_callback_sms==1) {
                            $("#before_callback_sms").prop("checked",true);
                        }
                    }


                }
        })
       
                    
    })

    
    

    $('.falcongrp-delete').click(function(e){
        var gridBatDel = dialog({
            title:'确认',
            width:250,
            content: '是否删除选中监控组',
            okValue:"确定",
            cancelValue:"取消",
            ok:function (){
                var tr = $(e.target).closest('tr');
                var Id = tr.find('.Id').val();
                $.post("/zjz_work_cmdb/cmdb/falcondelgrp",
                    {Id:Id}
                    ,function(result) {
                        // rere = $.parseJSON(result);
                        rere = result;
                        if(rere.success == false)
                        {
                            showWindows(rere.errInfo,'notice');
                            return;
                        }
                        else
                        {
                            showWindows("删除成功！",'success');
                            setTimeout(function(){window.location.reload();},1000);
                            return;
                        }
                    });

            },
            cancel: function () {
            }
        });
        gridBatDel.showModal();
    })


    
    $('.falcongrp-bindhost').click(function(e){
    
            $("#addhosts_div").hide('fast');
            $("#update_tpl").hide('fast');
            $("#bindtpl_div").hide('fast');
            $("#bind_tpl").hide('fast');
            $("#update").show('fast');



            var tr = $(e.target).closest('tr');
            var Id = tr.find('.Id').val();
            $("#grp_id").val(Id);
            $.ajax({
                url : '/zjz_work_cmdb/cmdb/falcongphost',
                method : 'post',
                data: {Id:Id},
                success : function(response){
                    console.log(response)
                     var output =  '<div class="panel panel-default" >';
                     output += '<div class="panel-heading">该监控组的监控设备</div>';
                     output += '<div class="panel-body">';
                     output += '<div class="pull-right">';
                     output += '<a class="btn btn-default" href="javascript:goto_hosts_add_div();">';
                     output += '<span class="glyphicon glyphicon-plus"></span></a>';
                     output += '</div>';
                     output += '<div style="line-height: 35px;">';
                     output += '监控设备列表';
                     output += '</div></div>'
               
            
                     output += '<table class="table table-hover table-bordered table-striped" style="margin-bottom: 0px;">';
                    
            
                     output += '<thead>';
                     output += '<tr>';
                     output += '<th>监控设备名称</th>';
                     output += '<th>维护起始时间</th>';
                     output += '<th>维护结束时间</th>';
                     output += '<th>关联监控组信息</th>';
                     output += '<th>关联模版信息</th>';
                     output += '<th>删除监控项</th>';
                     output += '<th>设置维护时间</th>';
                     
                     output += '</tr>';
                     output += '</thead>';
                     
                     
                    if(response.data==null) {
                         console.log(response.data)
                        
                     } else {
                        for (var i = 0; i < response.data.length; i ++ ) {
                       
                        output += '<tbody>';
                        output += '<tr>';
                        output += '<td>';
                        output += '<div class="hostname" title="dd_hostname">'+response.data[i][0].Hostname+'</div>'
                        output += '<input type="hidden" class="Id" value="'+response.data[i][0].Id+'">'
                        output += '</td>';
                        output += '<td>';
                        output += '<div class="begin_maintain" title="dd_begin_maintain">'+timeConverter(response.data[i][0].Maintain_begin)+'</div>'
                        output += '</td>';
                        output += '<td>';
                        output += '<div class="end_maintain" title="dd_end_maintain">'+timeConverter(response.data[i][0].Maintain_end)+'</div>'
                        output += '</td>';
                        output += '<td>';
                        output += '<a name="bindgroup" class="btn btn-success  btn-sm falconhg-bindgroup" >'+"设置"+'</a>';
                        output += '</td>';
                        output += '<td>';
                        output += '<a name="bindtemplate" class="btn btn-success  btn-sm falconhg-bindtemplate" >'+"设置"+'</a>'; 
                        output += '</td>';
                        output += '<td>';
                        output += '<a name="delhosts" class="btn btn-success  btn-sm falconhg-delhosts" href="javascript:hg_delhosts(\''+response.data[i][0].Id+'\');" >'+"删除"+'</a>'; 
                        output += '</td>';
                        output += '<td>';

                        output += '<div>'
                        output += '<input onclick="laydate({istime: true, format: \'YYYY-MM-DD hh:mm\'})" placeholder="begin" id="host_begin">'
                        
                        output += '<input onclick="laydate({istime: true, format: \'YYYY-MM-DD hh:mm\'})" placeholder="end" id="host_end">'
                        
                       
                        output += '<button class="btn btn-success btn-sm" onclick="host_maintain('+response.data[i][0].Id+')">维护</button>';
                        

                       
                        output += '<button class="btn btn-success btn-sm" onclick="no_host_maintain('+response.data[i][0].Id+')">取消维护</button>';
                        output += '</div>';

                        output += '</td>';
                        
                        output += '</tr>';
                        output += '</tbody>';
                        }
                     
                     }
                      

                    output += '</table>';
                    output += '</div>';
                   $('#update').html(output);
                    
            

                }
            });


    })
  
    

    $('.falcongrp-bindtpl').click(function(e){
    
            $("#addhosts_div").hide('fast');
            $("#update_tpl").hide('fast');
            $("#update").hide('fast');
            $("#bind_tpl").hide('fast');
            $("#bindtpl_div").show('fast');
            var tr = $(e.target).closest('tr');
            var Id = tr.find('.Id').val();
            $("#grp_id").val(Id);
            $.ajax({
                url : '/zjz_work_cmdb/cmdb/falcongptpl',
                method : 'post',
                data: {Id:Id},
                success : function(response){
                    console.log(response)
                     var output =  '<div class="panel panel-default" >';
                     output += '<div class="panel-heading">该监控组的关联模版</div>';
                     //output += '<div class="panel-body">';
                     //output += '<div class="pull-right">';
                     //output += '<a class="btn btn-default" href="javascript:goto_tpls_add_div('+Id+');">';
                    // output += '<span class="glyphicon glyphicon-plus"></span></a>';
                     //output += '</div>';
                     //output += '<div style="line-height: 35px;">';
                     //output += '关联模版列表';
                     //output += '</div></div>'
               
            
                     output += '<table class="table table-hover table-bordered table-striped" style="margin-bottom: 0px;">';
                    
            
                     output += '<thead>';
                     output += '<tr>';
                     output += '<th>关联模版名称</th>';
                     output += '<th>关联模版创建用户</th>';
                     output += '<th>取消关联模版</th>';
                    
                     
                     output += '</tr>';
                     output += '</thead>';
                     
                     
                    if(response.data==null) {
                         console.log(response.data)
                        
                     } else {
                        for (var i = 0; i < response.data.length; i ++ ) {
                            if (response.data[i]==null) {
                                console.log(response.data[i])
                            } else {
                                output += '<tbody>';
                        output += '<tr>';
                        output += '<td>';
                        output += '<div class="hostname" title="dd_hostname">'+response.data[i][0].Tpl_name+'</div>'
                        output += '<input type="hidden" class="Id" value="'+response.data[i][0].Id+'">'
                        output += '</td>';
                        output += '<td>';
                        output += '<div class="begin_maintain" title="dd_begin_maintain">'+response.data[i][0].Create_user+'</div>'
                        output += '</td>';
                        output += '<td>';
                        output += '<a name="deltpls" class="btn btn-success  btn-sm falconhg-delhosts" href="javascript:grp_deltpls(\''+response.data[i][0].Id+'\');" >'+"删除"+'</a>'; 
                        output += '</td>';
                        
                        output += '</tr>';
                        output += '</tbody>';
                       
                            }
                         }
                     
                     }
                    output += '</table>';
                    output += '</div>';
                   $('#update_tpl').html(output);
                }
            });
            $("#update_tpl").show('fast');
            


    })
  

    








    $('.falcontpl-delete').click(function(e){
        var gridBatDel = dialog({
            title:'确认',
            width:250,
            content: '是否删除选中模版',
            okValue:"确定",
            cancelValue:"取消",
            ok:function (){
                var tr = $(e.target).closest('tr');
                var Id = tr.find('.Id').val();
                $.post("/zjz_work_cmdb/cmdb/falcondeltpl",
                    {Id:Id}
                    ,function(result) {
                        // rere = $.parseJSON(result);
                        rere = result;
                        if(rere.success == false)
                        {
                            showWindows(rere.errInfo,'notice');
                            return;
                        }
                        else
                        {
                            showWindows("删除成功！",'success');
                            setTimeout(function(){window.location.reload();},1000);
                            return;
                        }
                    });

            },
            cancel: function () {
            }
        });
        gridBatDel.showModal();
    })


    $('.falcontpl-bind').click(function(e){
            $("#add_div").hide('fast');

            var tr = $(e.target).closest('tr');
            var Id = tr.find('.Id').val();
            $("#tpl_id").val(Id);
            $.ajax({
                url : '/zjz_work_cmdb/cmdb/falconshowgrptpl',
                method : 'post',
                data: {Id:Id},
                success : function(response){
                    console.log(response)

                    var output =  '<div class="panel panel-default" >';
                     output += '<div class="panel-heading">关联监控组</div>';
                     output += '<div class="panel-body">';
                     
                     //output += '<div class="pull-right">';
                     //output += '<a class="btn btn-default" href="javascript:goto_grps_add_div();">';
                     //output += '<span class="glyphicon glyphicon-plus"></span></a>';
                     //output += '</div>';

                     output += '<div style="line-height: 35px;">';
                     output += '关联监控组';
                     output += '</div></div>'
               
            
                     output += '<table class="table table-hover table-bordered table-striped" style="margin-bottom: 0px;">';
                    
            
                     output += '<thead>';
                     output += '<tr>';
                     output += '<th>监控组名称</th>';
                     output += '<th>监控组创建人</th>';
                     output += '<th>删除</th>';
                     
                     output += '</tr>';
                     output += '</thead>';
                     
                     
                    if(response.data==null) {
                         console.log(response.data)
                        
                     } else {
                        for (var i = 0; i < response.data.length; i ++ ) {
                        output += '<tbody>';
                        output += '<tr>';
                        output += '<td>';
                        output += '<div class="grpname" title="dd_grpname">'+response.data[i][0].Grp_name+'</div>'
                        output += '<input type="hidden" class="Id" value="'+response.data[i][0].Id+'">'
                        output += '</td>';
                        output += '<td>';
                        output += '<div class="create_user" title="dd_create_user">'+response.data[i][0].Create_user+'</div>'
                        output += '</td>';
                        output += '<td>';
                        output += '<a name="delgrps" class="btn btn-success  btn-sm falconhg-delgrps" href="javascript:tpl_delgrps(\''+response.data[i][0].Id+'\');" >'+"删除"+'</a>'; 
                        output += '</td>';
                        output += '</tr>';
                        output += '</tbody>';
                        }
                     
                     }
                      

                    output += '</table>';
                    output += '</div>';
                   $('#update_bind').html(output);

                    //cq
                }
            });
        })





	
	$('.falcontpl-add').click(function(e){
		
			//window.open("/zjz_work_cmdb/cmdb/falconnewstrategy")
			//window.location.href="/zjz_work_cmdb/cmdb/falconnewstrategy"
          $("#add_div").hide('fast');
          $("#bindparent_tpl_div").hide('fast');
          $("#uic_update").hide('fast');
          

	        var tr = $(e.target).closest('tr');
            var Id = tr.find('.Id').val();
            $("#tpl_id").val(Id);
            $.ajax({
                url : '/zjz_work_cmdb/cmdb/falconshowstrategy',
                method : 'post',
                data: {Id:Id},
                success : function(response){
                    console.log(response)
                   //console.log(response.data)

                     var output =  '<div class="panel panel-default" >';
                     output += '<div class="panel-heading">该模板中的策略列表</div>';
                     output += '<div class="panel-body">';
                     output += '<div class="pull-right">';
                     output += '<a class="btn btn-default" href="javascript:goto_strategy_add_div();">';
                     output += '<span class="glyphicon glyphicon-plus"></span></a>';
                     output += '</div>';
                     output += '<div style="line-height: 35px;">';
                     output += 'max: 最大报警次数 P：报警级别（&lt;3: 既发短信也发邮件 &gt;=3: 只发邮件） run：生效时间，不指定就是全天生效';
                     output += '</div></div>'
               
            
                     output += '<table class="table table-hover table-bordered table-striped" style="margin-bottom: 0px;">';
                     output += '<thead>';
                     output += '<tr>';
                     output += '<th><u>metric</u>/<span class="blue">tags</span><span class="gray">[note]</span></th>';
                     output += '<th>condition</th>';
                     output += '<th>max</th>';
                     output += '<th>P</th>';
                     output += '<th>run</th>';
                     output += '<th>修改</th>';
                     output += '<th>删除</th>';
                     output += '</tr>';
                     output += '</thead>';
                     
                     
                    if(response.data==null) {
                         console.log(response.data)
                        
                     } else {
                        for (var i = 0; i < response.data.length; i ++ ) {
                        output += '<tbody>';
                        output += '<tr>';
                        output += '<td>';
                        output += '<div class="mtn" title="dd_mtn">'+response.data[i].Metric+"/"+response.data[i].Tags+"["+response.data[i].Note+"]"+'</div>'
                        output += '</td>';
                        output += '<td>';
                        output += '<div class="condition" title="dd_condition">'+response.data[i].Func+response.data[i].Op+response.data[i].Right_vaue+'</div>'
                        output += '</td>';
                        output += '<td>';
                        output += '<div class="max" title="dd_max">'+response.data[i].Max_step+'</div>'
                        output += '</td>';
                        output += '<td>';
                        output += '<div class="P" title="dd_P">'+response.data[i].Priority+'</div>'
                        output += '</td>';
                        output += '<td>';
                        output += '<div class="run" title="dd_run">'+response.data[i].Run_begin+"-"+response.data[i].Run_end+'</div>'
                        output += '</td>';
                        output += '<td style="text-align: center;" role="gridcell">'
                        output += '<button type="button" onclick="falconstrategy_update('+response.data[i].Id+',\''+response.data[i].Metric+'\',\''+response.data[i].Tags+'\',\''+response.data[i].Note+'\',\''+response.data[i].Func+'\',\''+response.data[i].Op+'\',\''+response.data[i].Right_vaue+'\','+response.data[i].Max_step+','+response.data[i].Priority+',\''+response.data[i].Run_begin+'\',\''+response.data[i].Run_end+'\');" class="btn btn-success">修改策略</button>'
                        output += '</td>'
                        output += '<td style="text-align: center;" role="gridcell">'
                        output += '<button type="button" onclick="falconstrategy_del('+response.data[i].Id+');" class="btn btn-success">删除策略</button>'
                        output += '</td>'
                        
                        //output += '<td>';
                        //output += '<div class="operation" title="dd_operation">'+"operation"+'</div>'
                        //output += '</td>';
                        output += '</tr>';
                        output += '</tbody>';
                        }
                       
                      //  console.log(response.data.length)
                       
                      /*
                    <td>
                        <a href="javascript:clone_strategy('{{ s.id }}');" style="text-decoration: none;">
                            <span class="glyphicon glyphicon-duplicate orange"></span>
                        </a>
                        <span class="cut-line">¦</span>
                        <a href="javascript:modify_strategy('{{ s.id }}');" style="text-decoration: none;">
                            <span class="glyphicon glyphicon-edit orange"></span>
                        </a>
                        <span class="cut-line">¦</span>
                        <a href="javascript:delete_strategy('{{ s.id }}');" style="text-decoration: none;">
                            <span class="glyphicon glyphicon-trash orange"></span>
                        </a>
                    </td>
                    */









                     }
                      

                   /*
                   var output = '<table class="table table-striped table-bordered table-hover dataTables-example">';
                       output += '<colgroup>';
                       output += '<col style="width:15%"/>'
                       output += '<col style="width:15%"/>'
                       output += '<col style="width:25%"/>'
                       output += '<col style="width:15%"/>'
                       output += '<col style="width:15%"/>'
                       output += '<col style="width:15%"/>'
                       output += '</colgroup>'
                
                        output += '<thead>';
                        output += '<tr>';
                        output +='<th>主机名</th>';
                        output +='<th>业务IP</th>';
                        output +='<th>业务名称</th>';
                        output +='<th>模块名称</th>';
                        output +='<th>责任人</th>';
                        output +='<th>所属部门</th>';
                        output += '</tr>';
                        output += '</thead>';

                   for (var i = 0; i < response.data.length; i ++ ) {
                    if((response.data[i].Hostname.search(myExp) != -1) || (response.data[i].PublicIP.search(myExp) != -1) || (response.data[i].ApplicationName.search(myExp) != -1) || (response.data[i].ModuleName.search(myExp) != -1) || (response.data[i].Owner.search(myExp) != -1) || (response.data[i].Department.search(myExp) != -1)) {
                        
                        
                        output += '<tbody>';
                        output += '<tr>';
                        output += '<td>';
                        output += '<div class="host_name" title="dd_hostname">'+response.data[i].Hostname+'</div>'
                        output += '</td>';
                        output += '<td>';
                        output += '<div class="publicip" title="dd_publicip">'+response.data[i].PublicIP+'</div>'
                        output += '</td>';
                        output += '<td>';
                        output += '<div class="appname" title="dd_appname">'+response.data[i].ApplicationName+'</div>'
                        output += '</td>';
                        output += '<td>';
                        output += '<div class="modname" title="dd_modname">'+response.data[i].ModuleName+'</div>'
                        output += '</td>';
                        output += '<td>';
                        output += '<div class="owner" title="dd_owner">'+response.data[i].Owner+'</div>'
                        output += '</td>';
                        output += '<td>';
                        output += '<div class="department" title="dd_department">'+response.data[i].Department+'</div>'
                        output += '</td>';
                        output += '</tr>';
                        output += '</tbody>';

                       

                    }
                   };
                    output += '</table>'
                    */
                    output += '</table>';
                    output += '</div>';
                   $('#update').html(output);

                }
            });

        $("#update").show('fast');
    })


    $('.falconcls-delete').click(function(e){
        var gridBatDel = dialog({
            title:'确认',
            width:250,
            content: '是否删除选中自定义监控',
            okValue:"确定",
            cancelValue:"取消",
            ok:function (){
                var tr = $(e.target).closest('tr');
                var Id = tr.find('.Id').val();
                $.post("/zjz_work_cmdb/cmdb/falcondelcls",
                    {Id:Id}
                    ,function(result) {
                        // rere = $.parseJSON(result);
                        rere = result;
                        if(rere.success == false)
                        {
                            showWindows(rere.errInfo,'notice');
                            return;
                        }
                        else
                        {
                            showWindows("删除成功！",'success');
                            setTimeout(function(){window.location.reload();},1000);
                            return;
                        }
                    });

            },
            cancel: function () {
            }
        });
        gridBatDel.showModal();
    })

    $('.falconcls-edit').click(function(){
        var tr = $(this).closest('tr');
        $(this).hide();
        $(this).siblings().filter('a[name="saves"]').show();
        $(this).siblings().filter('a[name="cancels"]').show();

        var algorithm_value=tr.find('.algorithm').text();
        var endpoint_value=tr.find('.endpoint').text();
        var metric_value=tr.find('.metric').text();
        var tags_value=tr.find('.tags').text();


        tr.find('.algorithm').html('<input type="text" value="'+algorithm_value+'" style="width:100%;height:36px;"  class="k-input k-textbox algorithm_input">');
        tr.find(".algorithm .k-input").attr('placeholder','请输入集群算法');
        tr.find(".algorithm .k-input").attr('maxlength','32');

        tr.find('.endpoint').html('<input type="text" value="'+endpoint_value+'" style="width:100%;height:36px;"  class="k-input k-textbox endpoint_input">');
        tr.find(".endpoint .k-input").attr('placeholder','请输入集群监控设备名');
        tr.find(".endpoint .k-input").attr('maxlength','32');

        tr.find('.metric').html('<input type="text" value="'+metric_value+'" style="width:100%;height:36px;"  class="k-input k-textbox metric_input">');
        tr.find(".metric .k-input").attr('placeholder','请输入集群监控指标');
        tr.find(".metric .k-input").attr('maxlength','32');

        tr.find('.tags').html('<input type="text" value="'+tags_value+'" style="width:100%;height:36px;"  class="k-input k-textbox tags_input">');
        tr.find(".tags .k-input").attr('placeholder','请输入集群标签');
        tr.find(".tags .k-input").attr('maxlength','32');

    })


    $('.falconcls-cancel').click(function(){
        var tr = $(this).closest('tr');

        Algorithmname=tr.find('.Algorithm').val();
        Endpointname=tr.find('.Endpoint').val();
        Metricname=tr.find('.Metric').val();
        Tagsname=tr.find('.Tags').val();
        
        tr.find('.algorithm').html(Algorithmname);
        tr.find('.endpoint').html(Endpointname);
        tr.find('.metric').html(Metricname);
        tr.find('.tags').html(Tagsname);
        
        $(this).hide();
        $(this).siblings().filter('a[name="edits"]').show();
        $(this).siblings().filter('a[name="saves"]').hide();  
    })

    $('.falconcls-save').click(function(){
         
        var tr = $(this).closest('tr');
        var Id = tr.find('.Id').val();
            
        
        var nameDom1 = tr.find('.algorithm');
        var algorithmName = $.trim(nameDom1.find('input').val());

        var nameDom2 = tr.find('.endpoint');
        var endpointName = $.trim(nameDom2.find('input').val());

        var nameDom3 = tr.find('.metric');
        var metricName = $.trim(nameDom3.find('input').val());

        var nameDom4 = tr.find('.tags');
        var tagsName = $.trim(nameDom4.find('input').val());

        if((algorithmName.length> 50) || (algorithmName.length== 0)) {
            var diaCopyMsg = dialog({
                quickClose: true,
                align: 'left',
                padding:'5px 5px 5px 10px',
                skin: 'c-Popuplayer-remind-left',
                content: '<span style="color:#fff">请输入集群算法并限制在50位字符</span>'
            });
            diaCopyMsg.show(tr.find('.algorithm').find('input').get(0));
            return ;
        }


        if((endpointName.length> 16) || (endpointName.length== 0)) {
            var diaCopyMsg = dialog({
                quickClose: true,
                align: 'left',
                padding:'5px 5px 5px 10px',
                skin: 'c-Popuplayer-remind-left',
                content: '<span style="color:#fff">请输入集群监控设备名称并限制在16位字符</span>'
            });
            diaCopyMsg.show(tr.find('.endpoint').find('input').get(0));
            return ;
        }

        if((metricName.length> 16) || (metricName.length== 0)) {
            var diaCopyMsg = dialog({
                quickClose: true,
                align: 'left',
                padding:'5px 5px 5px 10px',
                skin: 'c-Popuplayer-remind-left',
                content: '<span style="color:#fff">请输入集群监控指标并限制在16位字符</span>'
            });
            diaCopyMsg.show(tr.find('.metric').find('input').get(0));
            return ;
        }

        if (tagsName.length> 16)  {
            var diaCopyMsg = dialog({
                quickClose: true,
                align: 'left',
                padding:'5px 5px 5px 10px',
                skin: 'c-Popuplayer-remind-left',
                content: '<span style="color:#fff">请输入集群监控标签并限制在16位字符</span>'
            });
            diaCopyMsg.show(tr.find('.tags').find('input').get(0));
            return ;
        }


        $.post("/zjz_work_cmdb/cmdbapp/falconeditcls",
            {
                id:Id,
                algorithm:algorithmName,
                endpoint:endpointName,
                metric:metricName,
                tags:tagsName
            }
            ,function(result) {
                //rere = $.parseJSON(result);
                rere = result;
                if(rere.success == false)
                {
                    showWindows(rere.errInfo,'notice');
                    return;
                }
                else
                {
                    showWindows('修改成功','success');
                    setTimeout(function(){window.location.reload();},1000);
                    return;
                    
                }
            });
    })
    

	function showWindows(Msg,level)
    {
    if(level=='success')
    {
        var d = dialog({
            width: 150,
            content: '<div class="c-dialogdiv2"><i class="c-dialogimg-success"></i>'+Msg+'</div>'
        });
        d.show();
    }
    else if(level =='error')
    {
        var d = dialog({
            title:'错误',
            width:300,
            okValue:"确定",
            ok:function(){},
            content: '<div class="c-dialogdiv2"><i class="c-dialogimg-failure"></i>'+Msg+'</div>'
        });
        d.showModal();
    }
    else
    {
        var d = dialog({
            title:'警告',
            width:300,
            okValue:"确定",
            ok:function(){},
            content: '<div class="c-dialogdiv2"><i class="c-dialogimg-prompt"></i>'+Msg+'</div>'
        });
        d.showModal();
    }
    }

});


      