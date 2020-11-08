//
//  filter.c
//  go2py
//
//  Created by PhuongNC on 4/30/20.
//  Copyright © 2020 PhuongNC. All rights reserved.
//

#include <Python.h>

/* Call C function */
char* call_func_filter_text(PyObject *func, const char* text)
{
    PyObject *args;
    PyObject *kwargs;
    PyObject *result = 0;
    char *retval;

    if (!PyEval_ThreadsInitialized())
    {
        PyEval_InitThreads();
        fprintf(stderr, "PyEval_InitThreads: is called.\n");
    }

    // Verify that func is a proper callable.
    if (!PyCallable_Check(func))
    {
        fprintf(stderr, "call_func_filter_text: expected a callable\n");
        goto fail;
    }

    // Convert a null-terminated C string to a Python object.
    args = Py_BuildValue("(s)", text);
    kwargs = NULL;

    // Call a callable Python object, with arguments given by the tuple args
    result = PyObject_Call(func, args, kwargs);
    Py_DECREF(args);
    Py_XDECREF(kwargs);

    // Occur an error?
    if (PyErr_Occurred())
    {
        PyErr_Print();
        goto fail;
    }

    // Return a pointer to the contents of 'result'
    retval = PyBytes_AsString(result);
    Py_DECREF(result);

    return retval;

    // Abort if happened error
    fail:
        Py_XDECREF(result);
        abort();
}

/* Load a symbol from a module */
PyObject *import_name(const char *modname, const char *symbol)
{
    PyObject *u_name, *module;
    u_name = PyUnicode_FromString(modname);
    module = PyImport_Import(u_name);
    PyObject *retval = PyObject_GetAttrString(module, symbol);

    Py_DECREF(u_name);
    Py_DECREF(module);

    return retval;
}

/* Load python environment */
void Py_Load() {
    Py_Initialize();
    PyRun_SimpleString("import sys");
    PyRun_SimpleString("sys.path.append('./')");
}

/* Unload python env*/
void Py_Unload() {
    Py_Finalize();
}

/* Filter text function */
char* Py_Filter_Text(const char *text) {
    PyObject * py_filter_func = import_name("filter", "filter_text");
    char *retval = call_func_filter_text(py_filter_func, text);
    Py_DECREF(py_filter_func);
    // Py_INCREF(retval);
    return retval;
}

/* MAIN */
#ifndef GO2PY_TEST
int main()
{
    return 0;
}
#endif
