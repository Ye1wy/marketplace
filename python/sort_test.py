import unittest
from sort_products import sort_price_rating


class TestingSort(unittest.TestCase):
    def test_okayness(self):
        a = {
            "sort": "rating",
            "products": [
                {
                    "template_id": "1",
                    "name": "Wireless Bluetooth Headphones",
                    "price": 59.99,
                    "rating": 4.5,
                    "link": "https://example.com/product/1"
                },
                {
                    "template_id": "2",
                    "name": "Smartphone Stand",
                    "price": 14.99,
                    "rating": 4.2,
                    "link": "https://example.com/product/2"
                },
                {
                    "template_id": "3",
                    "name": "4K Ultra HD TV",
                    "price": 499.99,
                    "rating": 4.8,
                    "link": "https://example.com/product/3"
                },
                {
                    "template_id": "4",
                    "name": "Laptop Backpack",
                    "price": 39.99,
                    "rating": 4.1,
                    "link": "https://example.com/product/4"
                },
                {
                    "template_id": "5",
                    "name": "Electric Toothbrush",
                    "price": 25.99,
                    "rating": 4.6,
                    "link": "https://example.com/product/5"
                },
                {
                    "template_id": "6",
                    "name": "Portable Bluetooth Speaker",
                    "price": 29.99,
                    "rating": 4.3,
                    "link": "https://example.com/product/6"
                },
                {
                    "template_id": "7",
                    "name": "Smart Fitness Tracker",
                    "price": 49.99,
                    "rating": 4.4,
                    "link": "https://example.com/product/7"
                },
                {
                    "template_id": "8",
                    "name": "Wireless Mouse",
                    "price": 19.99,
                    "rating": 4.0,
                    "link": "https://example.com/product/8"
                },
                {
                    "template_id": "9",
                    "name": "LED Desk Lamp",
                    "price": 24.99,
                    "rating": 4.2,
                    "link": "https://example.com/product/9"
                },
                {
                    "template_id": "10",
                    "name": "Portable Power Bank",
                    "price": 39.99,
                    "rating": 4.7,
                    "link": "https://example.com/product/10"
                },
                {
                    "template_id": "11",
                    "name": "Gaming Mouse",
                    "price": 49.99,
                    "rating": 4.6,
                    "link": "https://example.com/product/11"
                },
                {
                    "template_id": "12",
                    "name": "Noise Cancelling Headphones",
                    "price": 129.99,
                    "rating": 4.8,
                    "link": "https://example.com/product/12"
                },
                {
                    "template_id": "13",
                    "name": "Smartwatch",
                    "price": 99.99,
                    "rating": 4.3,
                    "link": "https://example.com/product/13"
                },
                {
                    "template_id": "14",
                    "name": "Electric Kettle",
                    "price": 29.99,
                    "rating": 4.4,
                    "link": "https://example.com/product/14"
                },
                {
                    "template_id": "15",
                    "name": "Cordless Vacuum Cleaner",
                    "price": 199.99,
                    "rating": 4.7,
                    "link": "https://example.com/product/15"
                },
                {
                    "template_id": "16",
                    "name": "Smart Home Security Camera",
                    "price": 69.99,
                    "rating": 4.5,
                    "link": "https://example.com/product/16"
                },
                {
                    "template_id": "17",
                    "name": "Digital Camera",
                    "price": 399.99,
                    "rating": 4.6,
                    "link": "https://example.com/product/17"
                },
                {
                    "template_id": "18",
                    "name": "Electric Grill",
                    "price": 59.99,
                    "rating": 4.3,
                    "link": "https://example.com/product/18"
                },
                {
                    "template_id": "19",
                    "name": "Smart Thermostat",
                    "price": 129.99,
                    "rating": 4.8,
                    "link": "https://example.com/product/19"
                },
                {
                    "template_id": "20",
                    "name": "Instant Pot",
                    "price": 89.99,
                    "rating": 4.7,
                    "link": "https://example.com/product/20"
                },
                {
                    "template_id": "21",
                    "name": "Smart Doorbell",
                    "price": 99.99,
                    "rating": 4.5,
                    "link": "https://example.com/product/21"
                },
                {
                    "template_id": "22",
                    "name": "Electric Skillet",
                    "price": 49.99,
                    "rating": 4.2,
                    "link": "https://example.com/product/22"
                },
                {
                    "template_id": "23",
                    "name": "Drone with Camera",
                    "price": 299.99,
                    "rating": 4.6,
                    "link": "https://example.com/product/23"
                },
                {
                    "template_id": "24",
                    "name": "Bluetooth Earbuds",
                    "price": 29.99,
                    "rating": 4.4,
                    "link": "https://example.com/product/24"
                },
                {
                    "template_id": "25",
                    "name": "Electric Pressure Cooker",
                    "price": 99.99,
                    "rating": 4.7,
                    "link": "https://example.com/product/25"
                },
                {
                    "template_id": "26",
                    "name": "Portable Mini Fridge",
                    "price": 79.99,
                    "rating": 4.3,
                    "link": "https://example.com/product/26"
                },
                {
                    "template_id": "27",
                    "name": "Air Purifier",
                    "price": 129.99,
                    "rating": 4.5,
                    "link": "https://example.com/product/27"
                },
                {
                    "template_id": "28",
                    "name": "Memory Foam Mattress",
                    "price": 499.99,
                    "rating": 4.8,
                    "link": "https://example.com/product/28"
                },
                {
                    "template_id": "29",
                    "name": "Smart Light Bulbs",
                    "price": 39.99,
                    "rating": 4.3,
                    "link": "https://example.com/product/29"
                },
                {
                    "template_id": "30",
                    "name": "Compact Air Conditioner",
                    "price": 199.99,
                    "rating": 4.6,
                    "link": "https://example.com/product/30"
                }
            ]
        }
        self.assertIsInstance(sort_price_rating(a), dict)

    def test_placeholder(self):
        self.assertTrue(True)

    def test_placeholder2(self):
        self.assertTrue(False)


if __name__ == '__main__':
    unittest.main()
