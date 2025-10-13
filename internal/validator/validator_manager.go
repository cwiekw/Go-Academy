package validator

type CharacterValidatorManager struct {
	validators map[uint64]CharacterValidator
}

func NewCharacterValidatorManager() *CharacterValidatorManager {
	return &CharacterValidatorManager{validators: make(map[uint64]CharacterValidator)}
}

func (v CharacterValidatorManager) AddValidator(movieId uint64, c CharacterValidator) {
	v.validators[movieId] = c
}

func (v CharacterValidatorManager) Validate(movieId uint64, name string) (bool, error) {
	validator, exists := v.validators[movieId]

	if !exists {
		return true, nil
	}

	return validator.Validate(name)
}
