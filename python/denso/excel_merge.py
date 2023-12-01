import pandas as pd

excel_files = [
    'files/01 NOVEMBER 2023 ADD.xlsx', 
    'files/01 NOVEMBER 2023.xlsx', 
    'files/02 NOVEMBER 2023.xlsx',
    'files/03 NOVEMBER 2023 ADD.xlsx'
]

merged_data = pd.DataFrame()

for file in excel_files:
    xls = pd.ExcelFile(file)
    print("xls-----------------------", xls)
    
    for sheet_name in xls.sheet_names:
        print('sheet_name-----------------------', sheet_name)
        df = pd.read_excel(file, sheet_name, skiprows=5)
        
        merged_data = pd.concat([merged_data, df], ignore_index=True)

merged_data.to_excel('merged_file.xlsx', index=False)
