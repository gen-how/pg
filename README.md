# Project Generator

以下是這個專案的執行計畫與步驟總結，以及幾個調用示例。
這個專案旨在創建一個名為 `pg` 的命令列工具，幫助用戶快速生成 C 和 C++ 專案的初始結構，並提供靈活的配置選項。

## 執行計畫與步驟

1. 專案目標
開發一個 CLI 工具 `pg`，用於生成 C 或 C++ 專案的初始結構。
支持生成應用程式專案、庫專案，或兩者兼具的混合專案。
支援多種建置系統（Makefile 和 CMake），允許用戶選擇一種或同時生成多種建置系統的配置文件。
提供許可證文件的生成，並允許用戶指定許可證類型。

2. 命令格式
基本格式：`pg <專案名稱> [選項]`
選項：
  -l 或 --lang：指定語言（c 或 cpp），此為必填項。
  --app：生成應用程式專案。
  --lib：生成庫專案。
  --make：生成 Makefile 配置文件。
  --cmake：生成 CMake 配置文件。
  --license：指定許可證類型（例如 mit、apache2 等），預設不產生。

3. 行為規則
專案類型：
  * 若未指定 --app 或 --lib，預設生成應用程式專案。
  * 可同時指定 --app 和 --lib，生成混合專案。
建置系統：
  * 可同時指定 --make 和 --cmake，生成多種建置系統的配置文件。
  * 若未指定建置系統，預設生成 Makefile。
許可證：
  * 若未指定 --license，預設不生成。
  * 用戶可指定其他許可證類型（如 apache2）。

4. 專案結構
應用程式專案：
  src/：包含 main.c 或 main.cpp。
  build/：用於存放編譯生成的文件。
  Makefile 和/或 CMakeLists.txt：根據用戶選擇生成。
  README.md：包含建置和運行說明。
  LICENSE：許可證文件（如果有指定 --license）。

庫專案：
  src/：包含庫的源碼文件（如 myproject.c 或 myproject.cpp）。
  include/：包含庫的頭文件（如 myproject.h）。
  build/：用於存放編譯生成的文件（僅在使用 Makefile 時）。
  Makefile 和/或 CMakeLists.txt：根據用戶選擇生成。
  README.md：包含建置和運行說明。
  LICENSE：許可證文件（如果有指定 --license）。

混合專案：
同時包含應用程式和庫的結構。
建置系統配置文件會處理應用程式和庫之間的依賴關係。

5. 錯誤處理
缺少專案名稱：以訊息提示。
未指定語言：以訊息提示。
無效的許可證類型：以訊息提示。

## 調用示例
### 示例 1：生成 C 語言應用程式專案，使用預設 Makefile
```bash
pg myproject -l c
```
生成結構：
```text
  myproject/
  ├── src/
  │   └── main.c
  ├── build/
  ├── Makefile
  └── README.md
```
說明：生成一個 C 語言應用程式專案，預設使用 Makefile。

### 示例 2：生成 C++ 庫專案，使用 CMake
```bash
pg myproject --lang=cpp --lib --cmake --license=mit
```

生成結構：
```text
  myproject/
  ├── src/
  │   └── myproject.cpp
  ├── include/
  │   └── myproject.h
  ├── CMakeLists.txt
  ├── README.md
  └── LICENSE
```
說明：生成一個 C++ 庫專案，使用 CMake 進行建置，使用 MIT 許可證。

### 示例 3：生成 C 語言混合專案，使用 Makefile 和 CMake
```bash
pg myproject --lang c --app --lib --make --cmake --license apache2
```
生成結構：
```text
  myproject/
  ├── src/
  │   ├── main.c
  │   └── myproject.c
  ├── include/
  │   └── myproject.h
  ├── build/
  ├── Makefile
  ├── CMakeLists.txt
  ├── README.md
  └── LICENSE
```
說明：生成一個 C 語言混合專案（應用程式 + 函式庫），同時生成 Makefile 和 CMake 配置文件並使用 Apache 2.0 許可證。

## 總結
這個執行計畫讓 `pg` 工具能夠靈活生成 C 和 C++ 專案結構，支持應用程式、庫或混合專案，並允許同時生成 Makefile 和 CMake 配置文件。
透過預設值（如 Makefile），簡化了常見使用場景，同時保留了用戶自定義選項的彈性，例如語言、建置系統和許可證類型。
