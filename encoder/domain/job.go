package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Job struct {
	ID               string    `json:"job_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	OutputBucketPath string    `json:"out_put_bucket_path" valid:"notnull"`
	Status           string    `json:"satus" valid:"notnull"`
	Video            *Video    `json:"video" valid:"-"`
	VideoID          string    `json:"_" valid:"-"  gorm:"type:uuid;column:video_id;notnull;foreignKey:videos (id);constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Error            string    `json:"error" valid:"-"`
	CreatedAt        time.Time `json:"created_at" valid:"-"`
	UpdateAt         time.Time `json:"update_at" valid:"-"`
}

func NewJob(output string, status string, video *Video) (*Job, error) {
	job := Job{
		OutputBucketPath: output,
		Status:           status,
		Video:            video,
	}

	job.prepare()

	err := job.Validate()
	if err != nil {
		return nil, err
	}
	return &job, nil
}

func (j *Job) Validate() error {
	_, err := govalidator.ValidateStruct(j)

	if err != nil {
		return err
	}
	return nil
}

func (j *Job) prepare() {
	j.ID = uuid.NewV4().String()
	j.CreatedAt = time.Now()
	j.UpdateAt = time.Now()
}
