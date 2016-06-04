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
		type: 'POST',
		host: '/',
		port: '9000',
		url: url,
		data: data,
		success: function (data, textStatus, jqXHR) {                
			var obj = JSON.parse(data)               
			console.dir(obj);
			$("#" + container).html(JSON.stringify(obj, null, 4));         
		},
		error: function (jqXHR, textStatus, errorThown) {
			console.error(jqXHR);
		}
	});
} 