function onDocumentReady() {
	console.log("testAPI start");
	renderTestTable();
};

var testMethodArr = [
	{
		method: "getRaceList(dt,city,name)",
		params: [
			'"dt":"2016-05-21"',
            '"city":"ара"',
            '"name":"рол"',
		],
		description: "возвращает массив гонок с базовыми характеристиками",
		func: `getRaceList()`,
		resultID: "getRaceList"
	},
	{
		method: "getRaceInfo(raceUID)",
		params: [
			'"raceUID":"edec71a9-f2cb-44c9-b142-d2d3747307b1"',
		],
		description: "возвращает полные данные по гонке",
		func: `getRaceInfo('edec71a9-f2cb-44c9-b142-d2d3747307b1')`,
		resultID: "getRaceInfo"
	},
	{
		method: "getClassList(raceUID)",
		params: [
			'"raceUID":"edec71a9-f2cb-44c9-b142-d2d3747307b1"',
		],
		description: "возвращает массив классов гонки",
		func: `getClassList('edec71a9-f2cb-44c9-b142-d2d3747307b1')`,
		resultID: "getClassList"
	},
	{
		method: "getClassInfo(raceUID,classUID)",
		params: [
			'"raceUID":"edec71a9-f2cb-44c9-b142-d2d3747307b1"',
			'"classUID":"ba37e1bc-9df6-4db8-884e-c68d127cec21"',
		],
		description: "возвращает данные по классу гонки",
		func: `getClassInfo('edec71a9-f2cb-44c9-b142-d2d3747307b1', 'ba37e1bc-9df6-4db8-884e-c68d127cec21')`,
		resultID: "getClassInfo"
	},
	{
		method: "getMarshalList(raceUID)",
		params: [
			'"raceUID":"edec71a9-f2cb-44c9-b142-d2d3747307b1"',
		],
		description: "возвращает массив маршалов гонки",
		func: `getMarshalList('edec71a9-f2cb-44c9-b142-d2d3747307b1')`,
		resultID: "getMarshalList"
	},
	{
		method: "getMarshalInfo(raceUID,mNumber)",
		params: [
			'"raceUID":"edec71a9-f2cb-44c9-b142-d2d3747307b1"',
			'"mNumber":2',
		],
		description: "возвращает данные по маршалу гонки",
		func: `getMarshalInfo('edec71a9-f2cb-44c9-b142-d2d3747307b1', 2)`,
		resultID: "getMarshalInfo"
	},
	{
		method: "getCheckpointList(raceUID,classUID)",
		params: [
			'"raceUID":"edec71a9-f2cb-44c9-b142-d2d3747307b1"',
			'"classUID":"ba37e1bc-9df6-4db8-884e-c68d127cec21"',
		],
		description: "возвращает массив контрольных точек по классу гонки",
		func: `getCheckpointList('edec71a9-f2cb-44c9-b142-d2d3747307b1', 'ba37e1bc-9df6-4db8-884e-c68d127cec21')`,
		resultID: "getCheckpointList"
	},
	{
		method: "getCheckpointInfo(raceUID,classUID,number)",
		params: [
			'"raceUID":"edec71a9-f2cb-44c9-b142-d2d3747307b1"',
			'"classUID":"ba37e1bc-9df6-4db8-884e-c68d127cec21"',
			'"number":2',
		],
		description: "возвращает данные по контрольной точке класса гонки",
		func: `getCheckpointInfo('edec71a9-f2cb-44c9-b142-d2d3747307b1', 'ba37e1bc-9df6-4db8-884e-c68d127cec21', 2)`,
		resultID: "getCheckpointInfo"
	},
];

function toggleResultView(id){
	$(id).toggle();
};

function renderTestTable() {
	console.log("renderTestTable");
	var obj = {};
	var testMethodRow = "";
	for (var i = 0; i < testMethodArr.length; i++) {
		obj = testMethodArr[i];
		method = obj.method;
		var params = "";
		for (var n = 0; n < obj.params.length; n++) {
			params = params + `<div>` + obj.params[n] + `</div>`;
		}
		description = obj.description;
		func = obj.func;
		resultID = obj.resultID;
		testMethodRow = `
			<tr>
				<td class="td1">` + obj.method + `</td>
				<td class="td2">` + params + `</td>
				<td class="td3">` + obj.description + `</td>
				<td class="td4">
					<button style="width:94px" class="btn btn-default" onclick="` + obj.func + `">go</button>
					<button style="width:94px" class="btn btn-default" onclick="toggleResultView('#` + obj.resultID + `')">результат</button>
				</td>
				<td class="td5">					
					<div id="` + obj.resultID + `" style="display: none;"></div>
				</td>
			</tr>  
		`;
		$("#table").append(testMethodRow);
	};
};

function getRaceList() {
	sendPost("/api/getRaceList", {
		"dt": "2016-05-21",
		"city": "ара",
		"name": "рол"
	}, "getRaceList");
};
function getRaceInfo(raceUID) {
	sendPost("/api/getRaceInfo", {
		"raceUID": raceUID,
	}, "getRaceInfo");
};
function getClassList(raceUID) {
	sendPost("/api/getClassList", {
		"raceUID": raceUID,
	}, "getClassList");
};
function getMarshalList(raceUID) {
	sendPost("/api/getMarshalList", {
		"raceUID": raceUID,
	}, "getMarshalList");
};
function getMarshalInfo(raceUID, mNumber) {
	sendPost("/api/getMarshalInfo", {
		"raceUID": raceUID,
		"mNumber": mNumber,
	}, "getMarshalInfo");
};
function getClassInfo(raceUID, classUID) {
	sendPost("/api/getClassInfo", {
		"raceUID": raceUID,
		"classUID": classUID,
	}, "getClassInfo");
};
function getCheckpointList(raceUID, classUID) {
	sendPost("/api/getCheckpointList", {
		"raceUID": raceUID,
		"classUID": classUID,
	}, "getCheckpointList");
};
function getCheckpointInfo(raceUID, classUID, num) {
	sendPost("/api/getCheckpointInfo", {
		"raceUID": raceUID,
		"classUID": classUID,
		"number": num,
	}, "getCheckpointInfo");
};

function sendPost(url, data, container) {
	$.ajax({
		type: "POST",
		host: "/",
		port: "9000",
		url: url,
		data: data,
		success: function (data, textStatus, jqXHR) {                
			var obj = JSON.parse(data)               
			console.dir(obj);
			$("#" + container).html(JSON.stringify(obj, null, 4));
			$("#" + container).parent().parent().addClass("success");
		},
		error: function (jqXHR, textStatus, errorThown) {
			console.error(jqXHR);
		}
	});
} 