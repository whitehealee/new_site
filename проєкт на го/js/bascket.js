/**
 * Created by Berezniuk.a on 05.06.2017.
 */
var d = document,
    itemBox = d.querySelectorAll('.item_box'), // блок каждого товара
    cartCont = d.getElementById('cart_content'); // блок вывода данных корзины
    cartBasket = d.getElementById('cart_basket');
// Функция кроссбраузерная установка обработчика событий
function addEvent(elem, type, handler){
    if(elem.addEventListener){
        elem.addEventListener(type, handler, false);
    } else {
        elem.attachEvent('on'+type, function(){ handler.call( elem ); });
    }
    return false;
}

// Получаем данные из LocalStorage
function getCartData(){
    return JSON.parse(localStorage.getItem('cart'));
}
// Записываем данные в LocalStorage
function setCartData(o){
    localStorage.setItem('cart', JSON.stringify(o));
    return false;
}
//Удаляем товар из корзины
function deleteToCart(itemId){
    this.disabled = true; // блокируем кнопку на время операции с корзиной
    var cartData = getCartData();
    console.log(JSON.stringify(cartData));

    console.log(itemId);
    if(cartData.hasOwnProperty(itemId)){ // если такой товар уже в корзине, то добавляем +1 к его количеству
        cartData[itemId][4] = false;
    }
    if(!setCartData(cartData)){
     this.disabled = false;
    }
    window.location.href = "/basket"
    // Обновляем данные в LocalStorage
    return false;
}
///////



// Добавляем товар в корзину
function addToCart(e){
    this.disabled = true; // блокируем кнопку на время операции с корзиной
    var cartData = getCartData() || {}, // получаем данные корзины или создаём новый объект, если данных еще нет
        parentBox = this.parentNode, // родительский элемент кнопки &quot;Добавить в корзину&quot;
        itemId = this.getAttribute('data-id'), // ID товара
        itemPhoto = parentBox.querySelector('.item_photo').innerHTML, //foto
        itemTitle = parentBox.querySelector('.item_title').innerHTML, // название товара
        itemPrice = parentBox.querySelector('.item_price').innerHTML; // стоимость товара
    if(cartData.hasOwnProperty(itemId)){ // если такой товар уже в корзине, то добавляем +1 к его количеству
        cartData[itemId][4] = true;
    } else { // если товара в корзине еще нет, то добавляем в объект
        cartData[itemId] = [itemPhoto, itemTitle, 1, itemPrice, true, itemId];
    }
    // Обновляем данные в LocalStorage
    if(!setCartData(cartData)){
        this.disabled = false; // разблокируем кнопку после обновления LS
        cartCont.innerHTML = 'Товар додано';
        cartBasket.innerHTML = 'Товар додано';
        setTimeout(function(){
            cartCont.innerHTML = 'Продовжити покупки...';
            cartBasket.innerHTML = 'Продовжити покупки...';
        }, 500);
    }
    return false;
}

// Устанавливаем обработчик события на каждую кнопку &quot;Добавить в корзину&quot;
for(var i = 0; i < itemBox.length; i++){
    addEvent(itemBox[i].querySelector('.add_item'), 'click', addToCart);
}

// Открываем корзину со списком добавленных товаров
function openCart(e){
    var cartData = getCartData(), // вытаскиваем все данные корзины
        totalItems = '';
    console.log(JSON.stringify(cartData));
    // если что-то в корзине уже есть, начинаем формировать данные для вывода
    if(cartData !== null){
        totalItems = '<form role="form" method="POST" action="/SaveOrder">'
        totalItems += '<table class="shopping_list"><tr><th></th><th></th><th></th><th></th></tr>';
        var count=0;
        for(var items in cartData){
            if(cartData[items][4]==true){
            totalItems += '<tr>';
            totalItems += '<td class="photoImageSize">' + '<img src="'+cartData[items][0]+'"+ alt="" class="photoImage">' + '</td>';
            for(var i = 1; i < cartData[items].length-2; i++){
                  switch (i) {
                  case 3:
                    totalItems += '<td>' + cartData[items][i] + ' грн/шт </td>';
                    break;
                  case 2:
                    totalItems += '<td>' + '<input type="number" class="kl_tovara" name="kl_tovara'+count+'" min="1" value="'+cartData[items][i]+'">'  + '</td>';
                    break;
                  default:
                    totalItems += '<td>' + cartData[items][i] + '</td>';
                }
            }
            totalItems += '<input type="text" class="itemID" id="itemID" name="itemID'+count+'" value="'+cartData[items][5]+'">';
            count++;
            totalItems += '<td>' + '<button id="delete_item" onclick="deleteToCart('+cartData[items][5]+')">&#4030;</button>' + '</td>';
            totalItems += '</tr>';
            }
        }
        count--;
        totalItems += '<input type="text" class="itemID"  name="count" value="'+count+'">';

        totalItems +=  '</table>';
        totalItems +=  ' <div class="basket_rigth">';
        totalItems +=  ' <p> <b>Для оформлення замовлення нам потрібні наступні дані:</b></p>';
        totalItems +=  ' <p> <input type="text" name="Lname" required oninvalid="this.setCustomValidity("Заповніть це поле!")" placeholder="Прізвище" value=""></p>';
        totalItems +=  ' <p> <input type="text" name="Fname" required oninvalid="this.setCustomValidity("Заповніть це поле!")" placeholder="Імя" value=""></p>';
        totalItems +=  ' <p> <input type="text" name="Oname" required oninvalid="this.setCustomValidity("Заповніть це поле!")" placeholder="По батькові" value=""></p>';
        totalItems +=  ' <p> <input type="text" name="Street"required oninvalid="this.setCustomValidity("Заповніть це поле!")" placeholder="Вулиця" value=""></p>';
        totalItems +=  ' <p> <input type="text" name="Home"  required oninvalid="this.setCustomValidity("Заповніть це поле!")" placeholder="№ дома" value=""></p>';
        totalItems +=  ' <p> <input type="text" name="Flat"           placeholder="№ квартири" value=""></p>';
        totalItems +=  ' <p> <input type="text" name="index" required oninvalid="this.setCustomValidity("Заповніть це поле!")" placeholder="Індекс" value=""></p>';
        totalItems +=  ' <p> <input type="text" name="Phone" required oninvalid="this.setCustomValidity("Заповніть це поле!")" placeholder="Телефон" value=""></p>';
        totalItems +=  ' <p><input type="submit" value="Замовити"></p>';
        totalItems +=  '</div>'
        totalItems +=  '</form>';
        totalItems +=  '<div class="clear"></div>';
        cartCont.innerHTML = totalItems;
        var item_count1 = 0;
        var item_count2 = 0;
        for(var items in cartData){
            item_count1++;
            if(cartData[items][4]==false){
                item_count2++;
            }
        }
        if(item_count1==item_count2){
            cartCont.innerHTML = 'Корзина порожня';
            cartBasket.innerHTML = 'Корзина порожня';
        }
    } else {
        // если в корзине пусто, то сигнализируем об этом
        cartCont.innerHTML = 'Корзина порожня';
        cartBasket.innerHTML = 'Корзина порожня';
    }
    return false;
}

/* Открыть корзину */
//addEvent(d.getElementById('checkout'), 'click', openCart);
document.addEventListener( "DOMContentLoaded", openCart);
/* Очистить корзину */
addEvent(d.getElementById('clear_cart'), 'click', function(e){
    localStorage.removeItem('cart');
    cartCont.innerHTML = 'Корзина порожня';
    cartBasket.innerHTML = 'Корзина порожня';
});

addEvent(d.getElementById('clear_cart'), 'click', function(e){
    localStorage.removeItem('cart');
    cartCont.innerHTML = 'Корзина порожня';
    cartBasket.innerHTML = 'Корзина порожня';
});
