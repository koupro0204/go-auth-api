import requests
import sys
from uuid import getnode as get_mac

HOST = "http://localhost:8080"
HEADERS = {'Content-Type': 'application/json'}
PRODUCT_NAME = "productname"
# macアドレスを取得（一意ではない可能性があるらしい）
PRODUCT_NUMBER = str(get_mac())

def create_product_request():
    url = f"{HOST}/products/create"
    payload = {"name": PRODUCT_NAME, "number": PRODUCT_NUMBER}
    
    response = requests.post(url, headers=HEADERS, json=payload)
    
    print(f"Status code: {response.status_code}")
    try:
        json_data = response.json()
        print(f"Error code: {json_data['code']}")
        print(f"Error message: {json_data['message']}")
    except ValueError:
        print("Product is created")

def create_user_request():
    url = f"{HOST}/user/create"
    email = input('Email: ')
    password = input('Password: ')
    payload = {"email": email, "password": password}
    
    response = requests.post(url, headers=HEADERS, json=payload)
    
    print(f"Status code: {response.status_code}")
    try:
        json_data = response.json()
        print(f"Error code: {json_data['code']}")
        print(f"Error message: {json_data['message']}")
    except ValueError:
        print("User is created")

def product_user_request():
    url = f"{HOST}/user/products"
    email = input('Email: ')
    password = input('Password: ')
    payload = {
        "email": email,
        "password": password,
        "name": PRODUCT_NAME,
        "number": PRODUCT_NUMBER
    }
    response = requests.post(url, headers=HEADERS, json=payload)
    print(f"Status code: {response.status_code}")
    try:
        json_data = response.json()
        print(f"Error code: {json_data['code']}")
        print(f"Error message: {json_data['message']}")
    except ValueError:
        print("User and product are linked")

def auth_request():
    url = f"{HOST}/user/auth"
    email = input('Email: ')
    password = input('Password: ')
    payload = {
        "email": email,
        "password": password,
        "name": PRODUCT_NAME,
        "number": PRODUCT_NUMBER
    }
    response = requests.post(url, headers=HEADERS, json=payload)
    print(f"Status code: {response.status_code}")
    try:
        json_data = response.json()
        if 'auth' in json_data:
            print("Auth is completed")
            return True
    except ValueError:
        pass 

    # エラー処理を行う
    if json_data['code']==4001:
        print('User and product are not linked')
        # お好みで
        print('do you link user and product?')
        product_user_request()
    elif json_data['code']==4002:
        print('user is not exist')
        # お好みで
        # print('do you create user?')
        # create_user_request()
    elif json_data['code']==4003:
        print('product is not exist')
        # お好みで
        # print('do you create product?')
        # create_product_request()
    elif json_data['code']==6002:
        print('password is incorrect')
    elif json_data['code']==6003:
        print('email is incorrect')

    print(f"error code: {json_data['code']}")
    print(f"error message: {json_data['message']}")
    return False

# メイン処理
def main():
    print("Hello World")


if __name__ == "__main__":
    # create_product_request()
    # create_user_request()
    # product_user_request()
    auth_bool = auth_request()
    if auth_bool==False:
        sys.exit( )
    main()
