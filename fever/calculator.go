package fever

import (
	"disco/base"
	"fmt"
)

func calculate(
	part_of_stack []*base.S,
	operater byte,
) (*base.S, error) {

	first := part_of_stack[0].Val

	switch first.(type) {
	case int64:
		answer, err := calculateInteger(part_of_stack, operater)
		if err != nil {
			return nil, err
		}

		return answer, nil

	case float64:
		answer, err := calculateFloat(part_of_stack, operater)
		if err != nil {
			return nil, err
		}

		return answer, nil

	default:
		return nil, fmt.Errorf("invalid argument type %s", first)
	}
}

func calculateFloat(
	part_of_stack []*base.S,
	operater byte,
) (*base.S, error) {

	answer := part_of_stack[0].Val.(float64)

	if operater == '+' {
		for _, s := range part_of_stack[1:] {
			num := s.Val

			switch num.(type) {
			case int64:
				answer += float64(num.(int64))

			case float64:
				answer += num.(float64)

			default:
				return nil, fmt.Errorf("invalid argument type %s", num)
			}
		}
	}

	if operater == '-' {
		for _, s := range part_of_stack[1:] {
			num := s.Val

			switch num.(type) {
			case int64:
				answer -= float64(num.(int64))

			case float64:
				answer -= num.(float64)

			default:
				return nil, fmt.Errorf("invalid argument type %s", num)
			}
		}
	}

	if operater == '*' {
		for _, s := range part_of_stack[1:] {
			num := s.Val

			switch num.(type) {
			case int64:
				answer *= float64(num.(int64))

			case float64:
				answer *= num.(float64)

			default:
				return nil, fmt.Errorf("invalid argument type %s", num)
			}
		}
	}

	if operater == '/' {
		for _, s := range part_of_stack[1:] {
			num := s.Val

			switch num.(type) {
			case int64:
				answer /= float64(num.(int64))

			case float64:
				answer /= num.(float64)

			default:
				return nil, fmt.Errorf("invalid argument type %s", num)
			}
		}
	}

	return base.MakeFloat(answer), nil
}

func calculateInteger(
	part_of_stack []*base.S,
	operater byte,
) (*base.S, error) {

	answer := part_of_stack[0].Val.(int64)

	if operater == '+' {
		for _, s := range part_of_stack[1:] {
			num := s.Val

			switch num.(type) {
			case int64:
				answer += num.(int64)

			case float64:
				answer += int64(num.(float64))

			default:
				return nil, fmt.Errorf("invalid argument type %s", num)
			}
		}

		return base.MakeInt(answer), nil
	}

	if operater == '-' {
		for _, s := range part_of_stack[1:] {
			num := s.Val

			switch num.(type) {
			case int64:
				answer -= num.(int64)

			case float64:
				answer -= int64(num.(float64))

			default:
				return nil, fmt.Errorf("invalid argument type %s", num)
			}
		}

		return base.MakeInt(answer), nil
	}

	if operater == '*' {
		for _, s := range part_of_stack[1:] {
			num := s.Val

			switch num.(type) {
			case int64:
				answer *= num.(int64)

			case float64:
				answer *= int64(num.(float64))

			default:
				return nil, fmt.Errorf("invalid argument type %s", num)
			}
		}

		return base.MakeInt(answer), nil
	}

	if operater == '/' {
		for _, s := range part_of_stack[1:] {
			num := s.Val

			switch num.(type) {
			case int64:
				answer /= num.(int64)

			case float64:
				answer /= int64(num.(float64))

			default:
				return nil, fmt.Errorf("invalid argument type %s", num)
			}
		}

		return base.MakeInt(answer), nil
	}

	if operater == '%' {
		for _, s := range part_of_stack[1:] {
			num := s.Val

			switch num.(type) {
			case int64:
				answer %= num.(int64)

			case float64:
				answer %= int64(num.(float64))

			default:
				return nil, fmt.Errorf("invalid argument type %s", num)
			}
		}

		return base.MakeInt(answer), nil
	}

	return nil, fmt.Errorf("calculate error")
}
