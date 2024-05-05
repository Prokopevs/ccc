package core

func (s *ServiceImpl) CheckCorectMultiplicatorType(str string) bool {
	found := false
	for _, v := range multiplicatorType {
		if v == str {
			found = true
			break
		}
	}

	return found
}


var multiplicatorType = []string{"gasStorage", "gasMining", "protection"} 
  