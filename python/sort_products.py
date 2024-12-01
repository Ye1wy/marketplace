import json
from typing import Any
import redis


def load_json(filename: str) -> dict[str, Any]:
    with open(filename, 'r', encoding='utf-8') as file:
        data = json.load(file)
    return data


def save_json(saved: dict, filename: str):
    with open(filename, 'w', encoding='UTF-8') as file:
        json.dump(saved, file, indent=4, ensure_ascii=False)


def normalize_rating(rating: float) -> float:
    return (rating - 1) / 4


def normalize_price(price: float, min_price: float, max_price: float) -> float:
    return (price - min_price) / (max_price - min_price)


def sort_price_rating(products: dict[str, Any]) -> dict[str, Any]:
    products_list = products['products']

    prices = [product['price'] for product in products['products']]
    min_price: float = min(prices)
    max_price: float = max(prices)

    sorted_products = {'request': products['request'], 'sort': products['sort'],
                       'products': sorted(products_list,
                                          key=lambda p: (0.4 * normalize_rating(p['rating']))
                                                        - (0.6 * normalize_price(p['price'],
                                                                                 min_price,
                                                                                 max_price)),
                                          reverse=True)}
    return sorted_products


def sort_price(products: dict[str, Any]) -> dict[str, Any]:
    sorted_products = {'request': products['request'], 'sort': products['sort'],
                       'products': sorted(products['products'],
                                          key=lambda p: p['price'],
                                          reverse=False)}
    return sorted_products


def sort_rating(products: dict[str, Any]) -> dict[str, Any]:
    sorted_products = {'request': products['request'], 'sort': products['sort'],
                       'products': sorted(products['products'],
                                          key=lambda p: p['rating'],
                                          reverse=True)}
    return sorted_products


def main():
    r = redis.Redis(host='localhost', port=6379, db=0)

    chanel_to_listen = 'scraper_to_sort'
    chanel_to_publish = 'sort_to_API'

    ps = r.pubsub()
    ps.subscribe(chanel_to_listen)

    for message in ps.listen():
        if message['type'] == 'message':
            try:
                data = json.loads(message['data'].decode('UTF-8'))
                if data['sort'] == 'default':
                    r.publish(chanel_to_publish, json.dumps(sort_price_rating(data)))
                elif data['sort'] == 'price':
                    r.publish(chanel_to_publish, json.dumps(sort_price_rating(data)))
                elif data['sort'] == 'rating':
                    r.publish(chanel_to_publish, json.dumps(sort_price_rating(data)))
                print('lol')
            except redis.RedisError as e:
                print(f'Could not push to redis: {e}')
    print('lol_kek')


if __name__ == '__main__':
    main()
