package flusher_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"ozonva/ova-task-api/internal/mocks"
	"ozonva/ova-task-api/internal/pkg/entities/tasks"
	"strconv"
	"time"

	"ozonva/ova-task-api/internal/flusher"
)

var _ = Describe("Flusher.Flush(chunkSize, repo)", func() {
	var (
		ctrl      *gomock.Controller
		mockRepo  *mocks.MockRepo
		chunkSize int
		f         flusher.Flusher
		db        []tasks.Task
	)
	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		f = nil
		mockRepo = nil
		chunkSize = 0
		db = make([]tasks.Task, 0, 0)
	})
	Context("when repo is nil", func() {
		It("flusher throw panic", func() {
			Expect(func() { flusher.New(2, nil) }).Should(PanicWith("entityRepo is nil"))
		})
	})
	Context("when repo working correct", func() {
		Context("when chunkSize less than one", func() {
			It("flusher throw panic", func() {
				Expect(func() { flusher.New(0, mockRepo) }).Should(PanicWith("chunkSize must be greater than 0"))
				Expect(func() { flusher.New(-1, mockRepo) }).Should(PanicWith("chunkSize must be greater than 0"))
			})
		})
		Context("when chunkSize greater than zero", func() {
			BeforeEach(func() {
				mockRepo = mocks.NewMockRepo(ctrl)
				mockRepo.EXPECT().AddTasks(gomock.Any()).DoAndReturn(func(entities []tasks.Task) error {
					db = append(db, entities...)
					return nil
				}).AnyTimes()
				chunkSize = 2
				f = flusher.New(chunkSize, mockRepo)
			})
			Context("when tasks is nil", func() {
				var notFlushed []tasks.Task
				BeforeEach(func() {
					notFlushed = f.Flush(nil)
				})
				It("flush do nothing", func() {
					Expect(db).Should(BeEmpty())
				})
				It("flush return empty slice", func() {
					Expect(notFlushed).Should(BeEmpty())
				})
			})
			Context("when tasks is empty", func() {
				var notFlushed []tasks.Task
				BeforeEach(func() {
					notFlushed = f.Flush([]tasks.Task{})
				})
				It("flush do nothing", func() {
					Expect(db).Should(BeEmpty())
				})
				It("flush return empty slice", func() {
					Expect(notFlushed).Should(BeEmpty())
				})
			})
			Context("when repo added successfully", func() {
				var notFlushed []tasks.Task
				var tasks = BuildFakeTasks(1, 2, 3, 4, 5)
				BeforeEach(func() {
					notFlushed = f.Flush(tasks)
				})
				It("flush add all", func() {
					Expect(db).Should(BeEquivalentTo(tasks))
				})
				It("flush return empty slice", func() {
					Expect(notFlushed).Should(BeEmpty())
				})
			})
		})
	})
	Context("when repo working with errors ", func() {
		var notFlushed []tasks.Task
		BeforeEach(func() {
			mockRepo = mocks.NewMockRepo(ctrl)
		})
		Context("when repo cant add nothing", func() {
			var entities = BuildFakeTasks(1, 2, 3, 4, 5)
			BeforeEach(func() {
				mockRepo.EXPECT().AddTasks(gomock.Any()).DoAndReturn(func(entities []tasks.Task) error {
					return errors.New("cant add entities")
				}).AnyTimes()
				chunkSize = 2
				f = flusher.New(chunkSize, mockRepo)
				notFlushed = f.Flush(BuildFakeTasks(1, 2, 3, 4, 5))
			})
			It("flush add nothing", func() {
				Expect(db).Should(BeEmpty())
			})
			It("flush return all tasks", func() {
				Expect(notFlushed).Should(BeEquivalentTo(entities))
			})
		})
		Context("when repo cant add just second chunk", func() {
			BeforeEach(func() {
				var callCounter = 0
				mockRepo.EXPECT().AddTasks(gomock.Any()).
					DoAndReturn(func(entities []tasks.Task) error {
						callCounter++
						if callCounter == 2 {
							db = append(db, entities...)
							return nil
						}
						return errors.New("cant add entities")
					}).AnyTimes()
				chunkSize = 2
				f = flusher.New(chunkSize, mockRepo)
				notFlushed = f.Flush(BuildFakeTasks(1, 2, 3, 4, 5))
			})
			It("flush just second chunk", func() {
				Expect(db).Should(BeEquivalentTo(BuildFakeTasks(3, 4)))
			})
			It("flush return all except second chunk", func() {
				Expect(notFlushed).Should(BeEquivalentTo(BuildFakeTasks(1, 2, 5)))
			})
		})
	})
})

var timeFakeTask, _ = time.Parse("2006.01.02", "2021.08.16")

func BuildFakeTasks(keys ...uint64) []tasks.Task {
	var result = make([]tasks.Task, 0, len(keys))
	for _, key := range keys {
		result = append(result, BuildFakeTask(key))
	}
	return result
}

func BuildFakeTask(key uint64) tasks.Task {
	return *tasks.New(key, key, strconv.Itoa(int(key)), timeFakeTask.Add(time.Duration(key)*time.Hour))
}
