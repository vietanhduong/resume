package home

import (
	"errors"
	"fmt"
)

// ValidateResume validate input resume
func ValidateResume(resume Resume) error {
	var err error
	// validate name field
	if err = ValidateRequiredField("name", resume.Name); err != nil {
		return err
	}
	// validate metadata block
	if err = ValidateMetadataBlock(resume.Metadata); err != nil {
		return err
	}
	// validate spec block
	if err = ValidateSpecBlock(resume.Spec); err != nil {
		return err
	}
	return nil
}

// ValidateMetadataBlock validate input metadata block
func ValidateMetadataBlock(meta Metadata) error {
	var err error
	// validate address
	if err = ValidateRequiredField("metadata.address", meta.Address); err != nil {
		return err
	}
	// validate phone
	if err = ValidateRequiredField("metadata.phone", meta.Phone); err != nil {
		return err
	}
	// validate email
	if err = ValidateRequiredField("metadata.email", meta.Email); err != nil {
		return err
	}
	// websites block is optional, no need to validate
	return nil
}

// ValidateSpecBlock validate input spec block
func ValidateSpecBlock(spec []Section) error {
	// validate spec block
	if spec == nil {
		return errors.New("spec is required")
	}
	// validate section blocks
	for i, section := range spec {
		if err := ValidateSectionBlock(fmt.Sprintf("spec.section[%d]", i), section); err != nil {
			return err
		}
	}
	return nil
}

// ValidateSectionBlock validate input section block
// with:
// 		father: indicates the parent block of block section, e.g: spec.section[1]
// 		achievement: input section block
func ValidateSectionBlock(father string, section Section) error {
	var err error
	// validate section name
	if err = ValidateRequiredField(fmt.Sprintf("%s.name", father), section.Name); err != nil {
		return err
	}
	// validate achievement blocks
	// achievement list cannot be empty
	if section.Achievements == nil {
		return errors.New(fmt.Sprintf("%s.achievements cannot be empty", father))
	}
	// validate each block
	for i, achievement := range section.Achievements {
		if err = ValidateAchievementBlock(fmt.Sprintf("%s.achievements[%d]", father, i), achievement); err != nil {
			return err
		}
	}
	return nil
}

// ValidateAchievementBlock validate input achievement block
// with:
// 		father: indicates the parent block of block achievement, e.g: spec.section[1].achievements[0]
// 		achievement: input achievement block
func ValidateAchievementBlock(father string, achievement Achievement) error {
	var err error
	// validate section name
	if err = ValidateRequiredField(fmt.Sprintf("%s.name", father), achievement.Name); err != nil {
		return err
	}
	// validate duration block
	if err = ValidateDurationBlock(fmt.Sprintf("%s.duration", father), achievement.Duration); err != nil {
		return err
	}
	// details block is optional, we can ignore them
	return nil
}

// ValidateDurationBlock validate input duration block
// with:
// 		father: indicates the parent block of block duration, e.g: spec.section[1].achievements[0].duration
// 		duration: input duration block
func ValidateDurationBlock(father string, duration Duration) error {
	// validate start field, this field is required
	if err := ValidateRequiredField(fmt.Sprintf("%s.start", father), duration.Start); err != nil {
		return err
	}
	// end field is optional
	return nil
}

// ValidateRequiredField validate input field, return error if input
// value is nil or empty.
// error format: <fieldName> is required
func ValidateRequiredField(fieldName, value string) error {
	if value == "" {
		return errors.New(fmt.Sprintf("%s is required", fieldName))
	}
	return nil
}
