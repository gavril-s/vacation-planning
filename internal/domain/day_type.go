package domain

type DayType string

const (
	// DayTypeUnknown - Нет информации
	DayTypeUnknown DayType = "unknown"
	// DayTypeWorkingDay - Рабочий день
	DayTypeWorkingDay DayType = "working_day"
	// DayTypeDayOff - Выходной день
	DayTypeDayOff DayType = "day_off"
	// DayTypeStateHoliday - Государственный праздник
	DayTypeStateHoliday DayType = "state_holiday"
	// DayTypeRegionalHoliday - Региональный праздник
	DayTypeRegionalHoliday DayType = "regional_holiday"
	// DayTypeShortenedDay - Предпраздничный сокращенный рабочий день
	DayTypeShortenedDay DayType = "shortened_day"
	// DayTypeAdditionalDayOff - Дополнительный / перенесенный выходной день
	DayTypeAdditionalDayOff DayType = "additional_day_off"
)
