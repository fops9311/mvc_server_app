asyncCallbacks.push(refresh_objects)
var obj_created = false
var odata
function refresh_objects(objects){
    if (!obj_created){
        let lp = transformIds(getObjectsIds(objects))
        odata = treedata(lp.labels, lp.parents, getObjectsCustomdata(objects))
        Plotly.newPlot('myDiv', odata, objectsLayout, { responsive: true });
        obj_created=true
    }else{
        let lp = transformIds(getObjectsIds(objects))
        Plotly.restyle('myDiv',
        {
            'labels':[lp.labels],
            'parents':[lp.parents],
            'customdata':[getObjectsCustomdata(objects)],
        })
        //Plotly.restyle('myDiv','labels',[lp.labels])
        //Plotly.restyle('myDiv','parents',[lp.parents])
        //Plotly.restyle('myDiv','customdata',[getObjectsCustomdata(objects)])
    }
}

var objectsLayout = {
    margin: {
        l: 5,
        r: 5,
        b: 15,
        t: 50,
        pad: 4
    },
}

function treedata(l, p, customdata) {
    l.push(appUser)
    p.push("")
        //customdata=customdata.sort((a, b) => a.length - b.length)
    customdata.push({ 0: "root", 1: 0 })
    return [{

        type: "treemap",

        branchvalues: "remainder",

        count: 'branches+leaves',

        maxdepth: 2,

        labels: l,

        parents: p,

        domain: { x: [0, 1] },

        textinfo: "label",

        labeltemplate: "'%{customdata[0]}'",

        texttemplate: "<i>%{customdata[0]}</i><br><b>(%{customdata[1]} Mpa)</b>",

        textposition: 'middle center',

        textfont: { size: 16 },

        outsidetextfont: { "size": 40, "color": "#377eb8" },

        marker: { 
            autocolorscale:false,
            "line": { "width": 1 }
         },

        customdata: customdata,


        //hovertemplate: '%{customdata[0]} last value %{customdata[1]}',

        pathbar: { "visible": false }

    }];
}


console.log("test transformObjects")
res = transformIds(["j1/3", "2/j1", "j5", "j6/2", "j1", "j1/3", "2", "j1", "j1", "j1"])
    //console.log(res.labels)
    //console.log(res.parents)

function getObjectsCustomdata(objects) {
    var result = []
    Object.entries(objects).sort((a, b) => a[1].Id.length - b[1].Id.length).map(o => {
        //console.log(o[1])
        //return
        var parts = o[1].Id.split("/")
        //parts.shift()
            //console.log(parts)
        result.push({
            0: parts[parts.length-1],
            1: o[1].Samples.pop().Value,
        })
    })
    return result
}

function getObjectsIds(objects) {
    var result = []
    objects.forEach((o) => {
        result.push(o.Id)
    })
    return result
}

function transformIds(Id) {
    var result = { labels: [], parents: [] }
    result.labels = Id.sort((a, b) => a.length - b.length)
    result.labels.every((v) => {
        parentCandidates(v).every((candidate) => {
            if (result.labels.includes(candidate)) {
                result.parents.push(candidate)
                return false
            }
            if (candidate === "") {
                result.parents.push(appUser)
                return false
            }
            return true
        })
        return true
    })
    console.log(result.labels)
    console.log(result.parents)
    return result
}

function parentCandidates(label) {
    dirs = label.split("/");
    if (dirs.length < 2) {
        return [""]
    }
    result = [];
    for (var i = 0; i < (dirs.length - 1); i++) {
        result[i] = dirs.slice(0, (dirs.length - 1) - i).join("/")
    }
    result.push("")
    return result
}