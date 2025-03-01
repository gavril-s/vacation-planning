package production_calendar

import (
	"fmt"
	"github.com/anatoliyfedorenko/isdayoff"
	"github.com/gavril-s/vacation-planning/internal/domain"
	"github.com/samber/lo"
	"time"
)

type Client struct {
	client *isdayoff.Client
}

func NewClient() *Client {
	return &Client{
		client: isdayoff.New(),
	}
}

func (c *Client) GetDayType(date domain.Date) (domain.DayType, error) {
	day, err := c.client.GetBy(isdayoff.Params{
		Year:        date.Year,
		Month:       lo.ToPtr(time.Month(date.Month)),
		Day:         lo.ToPtr(date.Day),
		CountryCode: lo.ToPtr(isdayoff.CountryCodeRussia),
		Pre:         lo.ToPtr(false),
		Covid:       lo.ToPtr(false),
	})
	if err != nil {
		return domain.DayTypeUnknown, fmt.Errorf("client.GetBy: %w", err)
	}
	if len(day) != 1 {
		return domain.DayTypeUnknown, fmt.Errorf("unexpected number of days returned: %d", len(day))
	}

	switch day[0] {
	case isdayoff.DayTypeWorking:
		return domain.DayTypeWorkingDay, nil
	case isdayoff.DayTypeNonWorking:
		return domain.DayTypeDayOff, nil
	default:
		return domain.DayTypeUnknown, fmt.Errorf("unexpected day type: %s", day[0])
	}
}
