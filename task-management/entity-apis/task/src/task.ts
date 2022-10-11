import express from 'express'
import dotenv from 'dotenv'
import prisma from './dbClient'
import { getReasonPhrase, StatusCodes } from 'http-status-codes';
import { Prisma } from '@prisma/client';

const taskRoute = express.Router();

taskRoute.get(`/`, async (req, res) => {
    const { taskGroupId } = req.query;

    const tasks = await prisma.task.findMany({
        where: {
            taskGroupId: Number(taskGroupId)
        },
        orderBy: {
            updatedAt: 'asc'
        }
    })

    res.status(StatusCodes.OK).json(tasks);
})

taskRoute.get(`/:id`, async (req, res) => {
    const { id } = req.params;
    const task = await prisma.task.findUnique({
        where: {
            id: Number(id),
        },
    })

    if (!task) {
        res.status(StatusCodes.NOT_FOUND).send(getReasonPhrase(StatusCodes.NOT_FOUND));
    } else {
        res.status(StatusCodes.OK).json(task);
    }
})

taskRoute.post(`/`, async (req, res) => {
    const { title, taskStatusId, taskGroupId } = req.body;
    const task = await prisma.task.create({
        data: {
            title: title,
            taskStatusId: taskStatusId,
            taskGroupId: taskGroupId
        }
    });
    res.status(StatusCodes.CREATED).json(task);
})

taskRoute.put(`/:id`, async (req, res) => {
    const { id } = req.params;
    const { title, taskStatusId, taskGroupId } = req.body;

    try {
        const task = await prisma.task.update({
            where: {
                id: Number(id)
            },
            data: {
                title: title,
                taskStatusId: taskStatusId,
                taskGroupId: taskGroupId
            }
        })
        res.status(StatusCodes.OK).json(task);
    } catch (e) {
        if (e instanceof Prisma.PrismaClientKnownRequestError) {
            console.log(e.message);
            res.status(StatusCodes.NOT_FOUND).send(getReasonPhrase(StatusCodes.NOT_FOUND));
        } else {
            throw e;
        }
    }

})

taskRoute.post(`/:id/status/:taskStatusId`, async (req, res) => {
    const { id, taskStatusId } = req.params;

    try {
        const task = await prisma.task.update({
            where: {
                id: Number(id)
            },
            data: {
                taskStatusId: Number(taskStatusId)
            }
        })
        res.status(StatusCodes.OK).json(task);
    } catch (e) {
        if (e instanceof Prisma.PrismaClientKnownRequestError) {
            console.log(e.message);
            res.status(StatusCodes.NOT_FOUND).send(getReasonPhrase(StatusCodes.NOT_FOUND));
        } else {
            throw e;
        }
    }

})

taskRoute.post(`/:id/group/:taskGroupId`, async (req, res) => {
    const { id, taskGroupId } = req.params;

    try {
        const task = await prisma.task.update({
            where: {
                id: Number(id)
            },
            data: {
                taskGroupId: Number(taskGroupId)
            }
        })
        res.status(StatusCodes.OK).json(task);
    } catch (e) {
        if (e instanceof Prisma.PrismaClientKnownRequestError) {
            console.log(e.message);
            res.status(StatusCodes.NOT_FOUND).send(getReasonPhrase(StatusCodes.NOT_FOUND));
        } else {
            throw e;
        }
    }

})


taskRoute.delete(`/:id`, async (req, res) => {
    const { id } = req.params

    try {
        const task = await prisma.task.delete({
            where: {
                id: Number(id),
            },
        })
        res.status(StatusCodes.OK).json(task);
    } catch (e) {
        if (e instanceof Prisma.PrismaClientKnownRequestError) {
            console.log(e.message);
            res.status(StatusCodes.NOT_FOUND).send(getReasonPhrase(StatusCodes.NOT_FOUND));
        } else {
            throw e;
        }
    }
})


export { taskRoute };   