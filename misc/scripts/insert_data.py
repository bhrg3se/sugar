import pandas as pd
import requests

BASE_URL = "http://localhost:8080/api/v1"
resp = requests.post(BASE_URL + "/auth/login", json={
    "email": "root",
    "password": "root",
})

token = resp.json()["body"]

doors_df = pd.read_csv("doors.csv")
#
# for i,door in doors_df.iterrows():
#     d={
#         "name":door["Door"],
#         "acmeDeviceID":door["Acme Device ID"],
#     }
#     cookies = {"token":token}
#     resp=requests.post(BASE_URL+"/admin/door/add",json=d,cookies=cookies)
#     print(resp.json())
#
#
#


# residents_df=pd.read_csv("residents.csv")
#
# for i,user in residents_df.iterrows():
#     d={
#         "email":user["Email"],
#         "password":"",
#         "first_name":user["First Name"],
#         "last_name":user["Last Name"],
#         "is_admin":False
#     }
#     cookies = {"token":token}
#     resp=requests.post(BASE_URL+"/admin/user/add",json=d,cookies=cookies)
#
