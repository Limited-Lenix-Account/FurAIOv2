import json

with open("util/misc/python_parse/japanese.json") as file:
    data = json.load(file)

    for i in data["searchRepository"]["pages"]["/japanese-whisky/b2b-158/?No=0&Nrpp=24&N=1441743281&Nr=B2CProduct.b2b_hasCase:Y"]["results"]["records"]:

        prod = i["attributes"]
        print(prod["product.displayName"][0], "$"+prod["product.listPrice"][0], prod["sku-B2CProduct.x_caseSize"][0] + " bottles / case")