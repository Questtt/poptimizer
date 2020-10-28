"""События связанные с таблицами."""
import dataclasses
import datetime

from poptimizer.data_di.shared.events import AbstractEvent


@dataclasses.dataclass(frozen=True)
class AppStarted(AbstractEvent):
    """Начало работы приложения.

    Обработчики данного события должны осуществить всю необходимую инициализацию.
    """

    timestamp: datetime.datetime = dataclasses.field(
        init=False,
        default_factory=datetime.datetime.utcnow,
    )


@dataclasses.dataclass(frozen=True)
class TradingDayEnded(AbstractEvent):
    """Произошло окончание очередного торгового дня."""

    date: datetime.date