"""Основные настраиваемые параметры"""
import logging
import pathlib

import pandas as pd


class POptimizerError(Exception):
    """Базовое исключение."""


# Конфигурация логгера
logging.basicConfig(level=logging.INFO)

# Количество колонок в распечатках без переноса на несколько страниц
pd.set_option("display.max_columns", 20)
pd.set_option("display.max_rows", 70)
pd.set_option("display.width", None)

# Путь к директории с данными
DATA_PATH = pathlib.Path(__file__).parents[1] / "data"

# Путь к директории с отчетам
REPORTS_PATH = pathlib.Path(__file__).parents[1] / "reports"

# Множитель, для переходя к после налоговым значениям
AFTER_TAX = 1 - 0.13

# Параметр для доверительных интервалов
T_SCORE = 2.0

# Максимальный объем одной торговой операции в долях портфеля
MAX_TRADE = 0.012

# Период в торговых днях, за который медианный оборот торгов
TURNOVER_PERIOD = 21 * 4

# Минимальный оборот - преимущества акции снижаются при приближении медианного оборота к данному уровню
TURNOVER_CUT_OFF = 3.9 * MAX_TRADE

# Параметры ML-модели
ML_PARAMS = {
    "data": (
        ("Label", {"days": 50, "on_off": True}),
        ("STD", {"days": 23, "on_off": True}),
        ("Ticker", {"on_off": True}),
        ("Mom12m", {"days": 259, "on_off": True, "periods": 2}),
        ("DivYield", {"days": 386, "on_off": True, "periods": 1}),
        ("Mom1m", {"days": 28, "on_off": False}),
        ("RetMax", {"days": 42, "on_off": True}),
        ("ChMom6m", {"days": 107, "on_off": True}),
    ),
    "model": {
        "bagging_temperature": 0.5549061439687702,
        "depth": 10,
        "l2_leaf_reg": 3.556206010583423,
        "learning_rate": 0.005016624020302364,
        "one_hot_max_size": 100,
        "random_strength": 0.5815581028215139,
    },
}
