package saver_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"ozonva/ova-task-api/internal/mocks"
	. "ozonva/ova-task-api/internal/pkg/entities/tasks"
	. "ozonva/ova-task-api/internal/saver"
	"ozonva/ova-task-api/internal/utils"
	"sync"
	"time"
)

var _ = Describe("Saver", func() {
	var (
		ctrl        *gomock.Controller
		mockFlusher *mocks.MockFlusher
		capacity    uint
		savePeriod  time.Duration
	)
	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		capacity = 0
		savePeriod = 0
		mockFlusher = nil
	})
	Context("when construct arguments are incorrect", func() {
		Context("when savePeriod <= 0", func() {
			It("saver throw panic", func() {
				mockFlusher = mocks.NewMockFlusher(ctrl)
				Expect(func() { NewSaver(1, mockFlusher, 0) }).Should(PanicWith("savePeriod must be greater than 0"))
				Expect(func() { NewSaver(1, mockFlusher, -1) }).Should(PanicWith("savePeriod must be greater than 0"))
			})
		})
		Context("when flusher is nil", func() {
			It("saver throw panic", func() {
				Expect(func() { NewSaver(1, nil, 1) }).Should(PanicWith("flusher is nil"))
			})
		})
	})
	Context("when construct arguments correct", func() {
		BeforeEach(func() {
			capacity = 2
			savePeriod = 1 * time.Nanosecond
			mockFlusher = mocks.NewMockFlusher(ctrl)
		})
		Context("when save tasks", func() {
			Context("when tasks flushed by period", func() {
				var (
					saver       Saver
					flushed     []Task = nil
					flushWaiter        = sync.WaitGroup{}
				)
				BeforeEach(func() {
					savePeriod = 1 * time.Nanosecond
					saver = NewSaver(capacity, mockFlusher, savePeriod)
					flushed = make([]Task, 0, 1)
					mockFlusher.EXPECT().Flush(gomock.Any()).DoAndReturn(func(entities []Task) []Task {
						defer flushWaiter.Done()
						flushed = append(flushed, entities...)
						return entities[:0]
					}).AnyTimes()
				})
				AfterEach(func() {
					saver.Close()
				})
				It("one task flushed", func() {
					task := Task{UserId: 1, TaskId: 1, Description: "task1", DateCreated: time.Time{}}
					flushWaiter.Add(1)

					saver.Save(task)
					isFlushed := utils.WaitTimeout(&flushWaiter, time.Second)

					Expect(isFlushed).Should(BeTrue())
					Expect(len(flushed)).Should(BeEquivalentTo(1))
					Expect(flushed[0]).Should(BeEquivalentTo(task))
				})
				It("many tasks flushed in different timeout", func() {
					task := Task{UserId: 1, TaskId: 1, Description: "task1", DateCreated: time.Time{}}

					tasks := 10
					for i := 0; i < tasks; i++ {
						flushWaiter.Add(1)
						saver.Save(task)
						isFlushed := utils.WaitTimeout(&flushWaiter, time.Second)
						Expect(isFlushed).Should(BeTrue())
					}

					Expect(len(flushed)).Should(BeEquivalentTo(tasks))
				})
			})
			Context("when tasks flushed on saver closing", func() {
				var (
					saver   Saver
					flushed []Task = nil
					flushWaiter        = sync.WaitGroup{}
				)
				BeforeEach(func() {
					savePeriod = 1 * time.Hour
					capacity = 10
					saver = NewSaver(capacity, mockFlusher, savePeriod)
					flushed = make([]Task, 0, 5)
					mockFlusher.EXPECT().Flush(gomock.Any()).DoAndReturn(func(entities []Task) []Task {
						flushed = entities
						flushWaiter.Done()
						return entities[:0]
					}).Times(1)
				})
				It("tasks flushed", func() {
					tasks := []Task{
						{UserId: 1, TaskId: 1, Description: "task1", DateCreated: time.Time{}},
						{UserId: 2, TaskId: 2, Description: "task2", DateCreated: time.Time{}},
						{UserId: 3, TaskId: 3, Description: "task3", DateCreated: time.Time{}},
						{UserId: 4, TaskId: 4, Description: "task4", DateCreated: time.Time{}},
						{UserId: 5, TaskId: 5, Description: "task5", DateCreated: time.Time{}},
					}
					flushWaiter.Add(len(tasks))
					for _, task := range tasks {
						saver.Save(task)
					}
					utils.WaitTimeout(&flushWaiter, 1 * time.Second)
					saver.Close()
					Expect(len(flushed)).Should(BeEquivalentTo(5))
					Expect(flushed).Should(BeEquivalentTo(tasks))
				})
			})
		})
	})
})