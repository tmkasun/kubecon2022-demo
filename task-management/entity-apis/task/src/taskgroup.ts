import express from 'express'
import prisma from './dbClient'
import { StatusCodes, getReasonPhrase } from 'http-status-codes';
import { Prisma } from '@prisma/client';

const taskGroupRoute = express.Router();

taskGroupRoute.get(`/`, async (req, res) => {
  const { userId } = req.query;

  const tasks = await prisma.taskGroup.findMany({
    where: {
      userId: {
        equals: userId as string
      }
    },
    orderBy: {
      updatedAt: 'asc'
    }
  })

  res.status(StatusCodes.OK).json(tasks);
})


taskGroupRoute.get(`/:id`, async (req, res) => {
  const { id } = req.params;
  const taskGroup = await prisma.taskGroup.findUnique({
    where: {
      id: Number(id),
    },
  })
  if (!taskGroup) {
    res.status(StatusCodes.NOT_FOUND).send(getReasonPhrase(StatusCodes.NOT_FOUND));
  } else {
    res.status(StatusCodes.OK).json(taskGroup);
  }
})

taskGroupRoute.post(`/`, async (req, res) => {
  const { title, userId } = req.body;
  const taskGroup = await prisma.taskGroup.create({
    data: {
      title: title,
      userId: userId
    }
  });
  res.status(StatusCodes.CREATED).json(taskGroup);
})

taskGroupRoute.put(`/:id`, async (req, res) => {
  const { id } = req.params;
  const { title } = req.body;

  try {
    const taskGroup = await prisma.taskGroup.update({
      where: {
        id: Number(id)
      },
      data: {
        title: title
      }
    })
    res.status(StatusCodes.OK).json(taskGroup);
  } catch (e) {
    if (e instanceof Prisma.PrismaClientKnownRequestError) {
      console.log(e.message);
      res.status(StatusCodes.NOT_FOUND).send(getReasonPhrase(StatusCodes.NOT_FOUND));
    } else {
      throw e;
    }
  }
})

taskGroupRoute.delete(`/:id`, async (req, res) => {
  const { id } = req.params

  try {
    const taskGroup = await prisma.taskGroup.delete({
      where: {
        id: Number(id),
      },
    })
    res.status(StatusCodes.OK).json(taskGroup);
  } catch (e) {
    if (e instanceof Prisma.PrismaClientKnownRequestError) {
      console.log(e.message);
      res.status(StatusCodes.NOT_FOUND).send(getReasonPhrase(StatusCodes.NOT_FOUND));
    } else {
      throw e;
    }

  }
})


export { taskGroupRoute };   