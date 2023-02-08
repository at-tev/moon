let dates_utc = document.getElementsByClassName('date');

for (let key in dates_utc) {
    if (key == "length") {
        break
    }

    let date = dates_utc[key].innerHTML;
    let dateUTC = new Date(date);
    dates_utc[key].innerHTML = dateUTC.toLocaleDateString();
};
