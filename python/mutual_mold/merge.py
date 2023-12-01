import pandas as pd
import os


def getExcelList():
    # 取得當前檔案的絕對路徑
    current_file_path = os.path.abspath(__file__)

    # 取得目前檔案所在的資料夾
    current_directory = os.path.dirname(current_file_path)

    # 組合資料夾路徑
    folder_path = os.path.join(current_directory, 'files')

    # 取得同層資料夾中的檔案列表
    files_in_folder = os.listdir(folder_path)

    # 列印所有excel檔
    print(files_in_folder)

    return files_in_folder



def mergeFiles():
    merged_data = pd.DataFrame()

    for file in getExcelList():
        xls = pd.ExcelFile(file)
        sheet_name = xls.sheet_names[1]  # 選擇第二個工作表
        df = pd.read_excel(file, sheet_name, skiprows=4)
        merged_data = pd.concat([merged_data, df], ignore_index=True)

    merged_data.to_excel('merged_file.xlsx', index=False)


mergeFiles()