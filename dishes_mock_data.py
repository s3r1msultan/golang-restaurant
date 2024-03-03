import json
import random
import faker

fake = faker.Faker()


def generate_real_dish_data(n):
    dishes_info = [
        {"name": "Pasta Carbonara",
         "img_URL": "https://images.pexels.com/photos/20315946/pexels-photo-20315946/free-photo-of-close-up-of-pasta-with-tomato-sauce-and-herbs-on-a-fork.jpeg?auto=compress&cs=tinysrgb&w=400"},
        {"name": "Margherita Pizza",
         "img_URL": "https://images.pexels.com/photos/11776376/pexels-photo-11776376.jpeg?auto=compress&cs=tinysrgb&w=400"},
        {"name": "Caesar Salad",
         "img_URL": "https://images.pexels.com/photos/2097090/pexels-photo-2097090.jpeg?auto=compress&cs=tinysrgb&w=400"},
        {"name": "Beef Bourguignon",
         "img_URL": "https://images.pexels.com/photos/15573468/pexels-photo-15573468/free-photo-of-bowl-with-spinach-chopped-beef-and-noodle.jpeg?auto=compress&cs=tinysrgb&w=400"},
        {"name": "Fish and Chips",
         "img_URL": "https://images.pexels.com/photos/13741669/pexels-photo-13741669.jpeg?auto=compress&cs=tinysrgb&w=400"},
        {"name": "Chicken Curry",
         "img_URL": "https://images.pexels.com/photos/7353487/pexels-photo-7353487.jpeg?auto=compress&cs=tinysrgb&w=400"},
        {"name": "Vegetable Stir Fry",
         "img_URL": "https://images.pexels.com/photos/3298693/pexels-photo-3298693.jpeg?auto=compress&cs=tinysrgb&w=400"},
        {"name": "Lamb Gyro", "img_URL": "https://images.pexels.com/photos/2871755/pexels-photo-2871755.jpeg?auto=compress&cs=tinysrgb&w=400"},
        {"name": "Quinoa Salad", "img_URL": "https://images.pexels.com/photos/566566/pexels-photo-566566.jpeg?auto=compress&cs=tinysrgb&w=400"},
        {"name": "Tofu Scramble", "img_URL": "https://images.pexels.com/photos/3026808/pexels-photo-3026808.jpeg?auto=compress&cs=tinysrgb&w=400"},
        {"name": "Mushroom Risotto",
         "img_URL": "https://images.pexels.com/photos/11190138/pexels-photo-11190138.jpeg?auto=compress&cs=tinysrgb&w=400"},
        {"name": "Tomato Soup", "img_URL": "https://images.pexels.com/photos/539451/pexels-photo-539451.jpeg?auto=compress&cs=tinysrgb&w=400"},
        {"name": "BBQ Ribs", "img_URL": "https://images.pexels.com/photos/16474897/pexels-photo-16474897/free-photo-of-a-plate-of-meat-barbecue-with-onions-and-pomegranate-seeds.jpeg?auto=compress&cs=tinysrgb&w=400"},
        {"name": "Falafel Wrap", "img_URL": "https://images.pexels.com/photos/8286779/pexels-photo-8286779.jpeg?auto=compress&cs=tinysrgb&w=400"},
        {"name": "Spinach Quiche", "img_URL": "https://images.pexels.com/photos/15573468/pexels-photo-15573468/free-photo-of-bowl-with-spinach-chopped-beef-and-noodle.jpeg?auto=compress&cs=tinysrgb&w=400"},
        {"name": "Duck Confit", "img_URL": "https://images.pexels.com/photos/3791089/pexels-photo-3791089.jpeg?auto=compress&cs=tinysrgb&w=400"}
    ]

    data = []
    for _ in range(n):
        dish = random.choice(dishes_info)
        name = dish["name"]
        img_url = dish["img_URL"]
        description = fake.text(max_nb_chars=200)
        price = float(round(random.uniform(5, 30), 2))
        weight = float(round(random.uniform(100, 500), 2))
        proteins = float(round(random.uniform(5, 30), 2))
        fats = float(round(random.uniform(5, 30), 2))
        carbohydrates = float(round(random.uniform(5, 30), 2))

        data.append({
            "name": name,
            "description": description,
            "price": price,
            "weight": weight,
            "protein": proteins,
            "fats": fats,
            "carbohydrates": carbohydrates,
            "img_URL": img_url
        })

    return data


dishes_data = generate_real_dish_data(50)
file_path = 'dishes_data.json'
with open(file_path, '+w') as file:
    json.dump(dishes_data, file, indent=4)
